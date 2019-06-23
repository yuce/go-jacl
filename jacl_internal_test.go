// Author: Yuce Tekol
// Created on: 2019-06-23, at: 07:50 +0300

package jacl

import (
	"errors"
	"testing"
)

func TestCountLeadingSpaces(t *testing.T) {
	compareCount(t, "", 0)
	compareCount(t, " ", 1)
	compareCount(t, "\t", 1)
	compareCount(t, "test 1", 0)
	compareCount(t, " test 2", 1)
	compareCount(t, "  test 3", 2)
	compareCount(t, "  test 4 ", 2)
	compareCount(t, "\ttest 5", 1)
	compareCount(t, "\t\ttest 6", 2)
	compareCount(t, "\t \ttest 7", 3)
	compareCount(t, "\t \tağaç", 3)
}

func TestTrimText(t *testing.T) {
	compareTrimmedText(t, "", "", nil)
	compareTrimmedText(t, " ", "", nil)
	compareTrimmedText(t, "\t", "", nil)
	compareTrimmedText(t, "\t \t  ", "", nil)
	compareTrimmedText(t, " test 1", "test 1", nil)
	compareTrimmedText(t, "\t test 2", "test 2", nil)
	compareTrimmedText(t, "\t\ntest 3", "test 3", nil)
	compareTrimmedText(t, "\t\ntest 4\nline2", "test 4\nline2", nil)
	compareTrimmedText(t, "\t\ntest 5\n line2", "test 5\n line2", nil)
	compareTrimmedText(t, "\t\ntest 6\n\nline2", "test 6\n\nline2", nil)
	compareTrimmedText(t, "\t\ntest 7\n\nline2", "test 7\n\nline2", nil)
	compareTrimmedText(t, "\t\ntest 8\n\nline2", "test 8\n\nline2", nil)
	compareTrimmedText(t, "\t\ntest 9\n\n  line2", "test 9\n\n  line2", nil)
	compareTrimmedText(t, "\t\n test 10\n\nline2", "", errors.New("inconsistent line start"))
}

func compareCount(t *testing.T, text string, target int) {
	count := countLeadingSpaces(text)
	if target != count {
		t.Fatalf("count text: %s\ncount %d != %d", text, target, count)
	}
}

func compareTrimmedText(t *testing.T, text string, target string, targetErr error) {
	trimmedText, err := trimText(text)
	if targetErr == nil {
		if err != nil {
			t.Fatalf("trim text: error nil != %s", err.Error())
		}
	} else {
		if err == nil {
			t.Fatalf("trim text: error %s != nil", targetErr.Error())
		} else if targetErr.Error() != err.Error() {
			t.Fatalf("trim text: error %s != %s", targetErr.Error(), err.Error())
		}
	}
	if target != trimmedText {
		t.Fatalf("trim text: %s\n%s\n!=\n%s", text, target, trimmedText)
	}
}
