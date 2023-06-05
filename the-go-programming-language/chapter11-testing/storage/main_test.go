package main

import (
	"strings"
	"testing"
)

// There is one problem: after this test function has returned, CheckQuota
// no longer works as it should because it's still using the test's fake implementation
// of notifyUser. (There is always a risk of this kind when updating global variables.)
// func TestCheckQuotaNotifyUser(t *testing.T) {
// 	var notifiedUser, notifiedMsg string
// 	notifyUser = func(user, msg string) {
// 		notifiedUser, notifiedMsg = user, msg
// 	}
// 	// ...simulate a 980MB-used condition...

// 	const user = "joe@example.org"
// 	CheckQuota(user)
// 	if notifiedUser == "" && notifiedMsg == "" {
// 		t.Fatalf("notifyUser not called")
// 	}
// 	if notifiedUser != user {
// 		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
// 	}
// 	const wantSubString = "98% of your quota"
// 	if !strings.Contains(notifiedMsg, wantSubString) {
// 		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubString)
// 	}
// }

// we must notify the test to restore the previous value so that subsequent
// test observe no effect, and we must do this on all execution pahts,
// including test failures and panics
// We can temporarily save and restore command-line flags, debugging options,
// and performance parameters
// Use global variable in this way is sofe only because go test does not normally runc
// multiple tests concurrently.
func TestCheckQuotaNotifyUser(t *testing.T) {
  // Save and restore original notifyUser
	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// ...simulate a 980MB-used condition...

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}
	const wantSubString = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubString) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubString)
	}
}
