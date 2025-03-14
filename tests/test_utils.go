/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at ApeCloud (https://www.apecloud.com/).
 * Copyright 2022-Present ApeCloud Co., Ltd
 */

package tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"time"

	ddtesting "github.com/DataDog/dd-sdk-go-testing"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"
	ddhttp "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/dnaeon/go-vcr.v3/cassette"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// RecordingMode defines valid usage of cassette recorder
type RecordingMode string

// Valid usage modes for cassette recorder
const (
	ModeIgnore    RecordingMode = "none"
	ModeReplaying RecordingMode = "false"
	ModeRecording RecordingMode = "true"
)

var testFiles2EndpointTags = map[string]map[string]string{}

// GetRecording returns the value of RECORD environment variable
func GetRecording() RecordingMode {
	if value, exists := os.LookupEnv("RECORD"); exists {
		switch value {
		case string(ModeIgnore):
			return ModeIgnore
		case string(ModeRecording):
			return ModeRecording
		default:
			return ModeReplaying
		}
	} else {
		return ModeReplaying
	}
}

// IsCIRun returns true if the CI environment variable is set to "true"
func IsCIRun() bool {
	return os.Getenv("CI") == "true"
}

// SecurePath replaces all dangerous characters in the path.
func SecurePath(path string) string {
	badChars := []string{"\\", "?", "%", "*", ":", "|", `"`, "<", ">", "'"}
	for _, c := range badChars {
		path = strings.ReplaceAll(path, c, "")
	}
	return filepath.Clean(path)
}

// Retry calls the call function for count times every interval while it returns false
func Retry(interval time.Duration, count int, call func() bool) error {
	for i := 0; i < count; i++ {
		if call() {
			return nil
		}
		if GetRecording() != ModeReplaying {
			time.Sleep(interval)
		}
	}
	return fmt.Errorf("Retry error: failed to satisfy the condition after %d times", count)
}

// ReadFixture opens the file at path and returns the contents as a string
func ReadFixture(path string) (string, error) {
	fixturePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get fixture file path: %v", err)
	}
	data, err := os.ReadFile(fixturePath)
	if err != nil {
		return "", fmt.Errorf("failed to open fixture file: %v", err)
	}
	return string(data), nil
}

// ConfigureTracer starts the tracer.
func ConfigureTracer(m *testing.M) {
	tracerOptions := make([]tracer.StartOption, 0, 2)
	if _, ok := os.LookupEnv("DD_SERVICE"); !ok {
		tracerOptions = append(tracerOptions, tracer.WithService("kb-cloud-client-go"))
	}
	if socketPath, ok := os.LookupEnv("DD_APM_RECEIVER_SOCKET"); ok {
		tracerOptions = append(tracerOptions, tracer.WithUDS(socketPath))
	}
	code := ddtesting.Run(m, tracerOptions...)
	os.Exit(code)
}

// getEndpointTagValue traverses callstack frames to find the test function that invoked this call;
// it then matches the file defining this function against testFiles2EndpointTags to figure out
// the tag value to set on span
func getEndpointTagValue(t *testing.T) (string, error) {
	var pcs [512]uintptr
	var frame runtime.Frame
	more := true
	n := runtime.Callers(1, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	functionFile := ""
	for more {
		frame, more = frames.Next()
		// nested test functions like `TestAuthenticationValidate/200_Valid` will have frame.Function ending with
		// ".funcX", `e.g. datadogV1.TestAuthenticationValidate.func1`, so trim everything after last "/" in test name
		// and everything after last "." in frame function name
		frameFunction := frame.Function
		testName := t.Name()
		if strings.Contains(testName, "/") {
			testName = testName[:strings.LastIndex(testName, "/")]
			frameFunction = frameFunction[:strings.LastIndex(frameFunction, ".")]
		}
		if strings.HasSuffix(frameFunction, "."+testName) {
			functionFile = frame.File
			// when we find the frame with the current test function, match it against testFiles2EndpointTags
			for subdir, file2tag := range testFiles2EndpointTags {
				for file, tag := range file2tag {
					if strings.HasSuffix(frame.File, fmt.Sprintf("%s/%s.go", subdir, file)) {
						return tag, nil
					}
				}
			}
		}
	}
	return "", fmt.Errorf(
		"endpoint tag for test file %s not found in tests/test_utils.go, please add it to `testFiles2EndpointTags`",
		functionFile)
}

// WithTestSpan starts new span with test information.
func WithTestSpan(ctx context.Context, t *testing.T) (context.Context, func()) {
	t.Helper()
	tag, err := getEndpointTagValue(t)
	if err != nil {
		t.Log(err.Error())
		tag = "features"
	}
	ctx, finish := ddtesting.StartTestWithContext(ctx, t, ddtesting.WithSkipFrames(2), ddtesting.WithSpanOptions(
		// Set resource name to TestName
		tracer.ResourceName(t.Name()),

		// We need to make the tag be something that is then searchable in monitors
		// https://docs.datadoghq.com/tracing/guide/metrics_namespace/#errors
		// "version" is really the only one we can use here
		// NOTE: version is treated in slightly different way, because it's a special tag;
		// if we set it in StartSpanFromContext, it would get overwritten
		tracer.Tag(ext.Version, tag),
	))

	return ctx, func() {
		if r := recover(); r != nil {
			t.Errorf("test paniced: %v", r)
		}
		finish()
	}
}

func createWithDir(path string) (*os.File, error) {
	dirName := filepath.Dir(path)
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.Create(path)
}

// SetClock stores current time in .freeze file.
func SetClock(path string) (clockwork.FakeClock, error) {
	now := clockwork.NewRealClock().Now()
	if GetRecording() == ModeRecording {
		os.MkdirAll("cassettes", 0755)

		f, err := createWithDir(fmt.Sprintf("cassettes/%s.freeze", path))
		if err != nil {
			return nil, err
		}
		defer f.Close()
		f.WriteString(now.Format(time.RFC3339Nano))
	}
	return clockwork.NewFakeClockAt(now), nil
}

// RestoreClock restore current time from .freeze file.
func RestoreClock(path string) (clockwork.FakeClock, error) {
	freezePath := fmt.Sprintf("cassettes/%s.freeze", path)
	data, err := os.ReadFile(freezePath)
	if err != nil {
		return nil, fmt.Errorf("time file '%s' not found: create one setting `RECORD=true` or ignore it using `RECORD=none`", freezePath)
	}
	now, err := time.Parse(time.RFC3339Nano, string(data))
	if err != nil {
		return nil, err
	}
	return clockwork.NewFakeClockAt(now), nil
}

type contextKey string

var (
	clockKey = contextKey("clock")
)

// WithClock sets clock to context.
func WithClock(ctx context.Context, path string) (context.Context, error) {
	var fc clockwork.FakeClock
	var err error
	if GetRecording() != ModeReplaying {
		fc, err = SetClock(path)
	} else {
		fc, err = RestoreClock(path)
	}
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, clockKey, fc), nil
}

// UniqueEntityName will return a unique string that can be used as a title/description/summary/...
// of an API entity.
func UniqueEntityName(ctx context.Context, t *testing.T) *string {
	name := WithUniqueSurrounding(ctx, t.Name())
	return &name
}

// WithUniqueSurrounding will wrap a string that can be used as a title/description/summary/...
// of an API entity.
func WithUniqueSurrounding(ctx context.Context, name string) string {
	prefix := "Test"
	if GetRecording() == ModeIgnore {
		// In ignore mode we add the language prefix to track unremoved data
		prefix = "Test-Go"
	}
	alnum := regexp.MustCompile(`[^A-Za-z0-9]+`)

	name = string(alnum.ReplaceAll([]byte(name), []byte("_")))
	maxSize := len(name)
	if maxSize > 100 {
		maxSize = 100
	}

	// NOTE: some endpoints have limits on certain fields (e.g. Roles V2 names can only be 55 chars long),
	// so we need to keep this short
	result := fmt.Sprintf("%s-%s-%d", prefix, name[:maxSize], ClockFromContext(ctx).Now().Unix())
	// In case this is used in URL, make sure we replace the slash that is added by subtests
	result = strings.ReplaceAll(result, "/", "-")
	return result
}

// ClockFromContext returns clock or panics.
func ClockFromContext(ctx context.Context) clockwork.FakeClock {
	if ctx == nil {
		log.Fatal("ctx is required")
	}
	v := ctx.Value(clockKey)
	fc, ok := v.(clockwork.FakeClock)
	if !ok {
		log.Fatalf("invalid value %v should be clockwork.FakeClock{}", v)
	}
	return fc
}

func removeURLSecrets(u *url.URL) string {
	q := u.Query()
	q.Del("api_key")
	q.Del("application_key")
	u.RawQuery = q.Encode()
	site, ok := os.LookupEnv("KB_CLOUD_TEST_SITE")
	if ok {
		u.Host = strings.Replace(u.Host, site, "apecloud.cn", 1)
	}
	return u.String()
}

// CompareAsJSON returns true if JSON strings serialize into same values.
func CompareAsJSON(first, second io.Reader) (bool, error) {
	var f, s interface{}
	d := common.NewDecoder(first)
	if err := d.Decode(&f); err != nil {
		return false, err
	}
	d = common.NewDecoder(second)
	if err := d.Decode(&s); err != nil {
		return false, err
	}
	return reflect.DeepEqual(f, s), nil
}

// MatchInteraction checks if the request matches a store request in the given cassette.
func MatchInteraction(r *http.Request, i cassette.Request) bool {
	// Default matching on method
	if r.Method != i.Method {
		return false
	}
	cassetteURL, err := url.Parse(i.URL)
	if err != nil {
		return false
	}

	for k, v := range i.Headers {
		if strings.Compare(k, "Content-Type") != 0 {
			// FIXME enable when Accept headers are correctly generated
			// if strings.Compare(k, "Accept") != 0 && strings.Compare(k, "Content-Type") != 0 {
			continue
		}

		if strings.Compare(k, "Content-Type") == 0 && strings.HasPrefix(i.Headers.Get(k), "multipart/form-data") {
			continue
		}

		rv := r.Header.Values(k)
		if len(v) != len(rv) {
			return false
		}
		for kv, vv := range rv {
			if strings.Compare(vv, v[kv]) != 0 {
				log.Printf("header %s %s does not match %s on position %d", k, v, rv, kv)
				return false
			}
		}
	}

	if cassetteURL.Path != r.URL.Path {
		return false
	}

	q := r.URL.Query()
	q.Del("api_key")
	q.Del("application_key")
	for k, v := range cassetteURL.Query() {
		vq, ok := q[k]
		if !ok {
			log.Printf("query param %s is missing", k)
			return false
		}
		for kv, vv := range v {
			if strings.Compare(vv, vq[kv]) != 0 {
				vvnow, verr := time.Parse(time.RFC3339Nano, vv)
				vqnow, vverr := time.Parse(time.RFC3339Nano, vq[kv])
				if verr == nil && vverr == nil && vvnow.Unix() == vqnow.Unix() {
					continue
				}
				log.Printf("query param %s does not match %s stored value %s", k, vv, vq[kv])
				return false
			}
		}
	}

	// Request does not contain body (e.g. `GET`)
	if r.Body == nil {
		return i.Body == ""
	}

	// Load request body
	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		return false
	}
	r.Body = io.NopCloser(&b)

	matched := (b.String() == "" || b.String() == i.Body)

	// Ignore boundary differences for multipart/form-data content
	if !matched && strings.HasPrefix(r.Header["Content-Type"][0], "multipart/form-data") {
		rl := strings.Split(strings.TrimSpace(b.String()), "\n")
		cl := strings.Split(strings.TrimSpace(i.Body), "\n")
		if len(rl) > 1 && len(cl) > 1 {
			rs := strings.Join(rl[1:len(rl)-1], "\n")
			cs := strings.Join(cl[1:len(cl)-1], "\n")
			if rs == cs {
				matched = true
			}
		}
	} else if !matched {
		matched, err = CompareAsJSON(strings.NewReader(b.String()), strings.NewReader(i.Body))
		if err != nil {
			log.Fatalf("failed to compare as JSON: %s", err)
			return false
		}
	}
	return matched
}

// FilterInteraction removes secret arguments from the URL.
func FilterInteraction(i *cassette.Interaction) error {
	u, err := url.Parse(i.Request.URL)
	if err != nil {
		return err
	}
	i.Request.URL = removeURLSecrets(u)
	i.Request.Headers.Del("Dd-Api-Key")
	i.Request.Headers.Del("Dd-Application-Key")
	// Remove custom URL in report-uri from response
	i.Response.Headers.Del("Content-Security-Policy")
	return nil
}

// Recorder intercepts HTTP requests.
func Recorder(ctx context.Context, name string) (*recorder.Recorder, error) {
	// Configure recorder
	var mode recorder.Mode
	switch GetRecording() {
	case ModeReplaying:
		mode = recorder.ModeReplayOnly
	case ModeIgnore:
		mode = recorder.ModePassthrough
	default:
		mode = recorder.ModeRecordOnly
	}

	opts := &recorder.Options{
		CassetteName:       fmt.Sprintf("cassettes/%s", name),
		Mode:               mode,
		SkipRequestLatency: true,
		RealTransport:      http.DefaultTransport,
	}

	rec, err := recorder.NewWithOptions(opts)
	if err != nil {
		return nil, err
	}

	if span, ok := tracer.SpanFromContext(ctx); ok {
		span.SetTag("recorder.mode", mode)
	}

	rec.SetMatcher(MatchInteraction)
	rec.AddHook(FilterInteraction, recorder.AfterCaptureHook)

	return rec, nil
}

// WrapRoundTripper includes tracing information.
func WrapRoundTripper(rt http.RoundTripper, opts ...ddhttp.RoundTripperOption) http.RoundTripper {
	if GetRecording() == ModeReplaying {
		return rt
	}
	return ddhttp.WrapRoundTripper(
		rt,
		ddhttp.WithAfter(func(r *http.Response, span ddtrace.Span) {
			if 500 <= r.StatusCode && r.StatusCode < 600 {
				var b bytes.Buffer
				tee := io.TeeReader(r.Body, &b)
				msg, _ := io.ReadAll(tee)

				span.SetTag(ext.Error, true)
				span.SetTag(ext.ErrorMsg, msg)
				span.SetTag(ext.ErrorDetails, r.Status)

				r.Body = io.NopCloser(&b)
			}
		}),
	)
}

// Assertions wrapper
type Assertions struct {
	require.Assertions
}

// TestingT keeps track of a context.
type TestingT struct {
	*testing.T
	ctx context.Context
}

// Errorf format error message.
func (t *TestingT) Errorf(format string, args ...interface{}) {
	t.T.Helper()
	span, _ := tracer.StartSpanFromContext(t.ctx, "testing.error")
	defer span.Finish()
	span.SetTag(ext.Error, fmt.Errorf(format, args...))
	t.T.Errorf(format, args...)
}

// Assert wraps context and testing object.
func Assert(ctx context.Context, t *testing.T) *Assertions {
	t.Helper()
	return &Assertions{*require.New(&TestingT{t, ctx})}
}
