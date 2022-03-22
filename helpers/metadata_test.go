package helpers

import (
	tektonv1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"testing"
)

var (
	testObject      = &tektonv1beta1.PipelineRun{}
	testAnnotations = map[string]string{
		"a.test/annotation":              "foo",
		"another.test/annotation-string": "bar",
	}
	testLabels = map[string]string{
		"a.test/label":              "foo",
		"another.test/label-string": "bar",
	}
)

func init() {
	testObject.SetAnnotations(testAnnotations)
	testObject.SetLabels(testLabels)
}

func TestHasAnnotation(t *testing.T) {
	checkExistingAnnotation := HasAnnotation(testObject, "a.test/annotation")
	if checkExistingAnnotation != true {
		t.Fatalf("HasAnnotation() failed to find existing annotation")
	}

	checkMissingAnnotation := HasAnnotation(testObject, "bar")
	if checkMissingAnnotation != false {
		t.Fatalf("HasAnnotation() found a non-existent annotation")
	}
}

func TestHasAnnotationWithValue(t *testing.T) {
	checkAnnotationProperValue := HasAnnotationWithValue(testObject, "a.test/annotation", "foo")
	if checkAnnotationProperValue != true {
		t.Fatalf("HasAnnotationWithValue() failed to find value of existing annotation")
	}

	checkAnnotationWrongValue := HasAnnotationWithValue(testObject, "a.test/annotation", "bar")
	if checkAnnotationWrongValue != false {
		t.Fatalf("HasAnnotationWithValue() found the wrong value of an existing annotation")
	}

	checkMissingAnnotation := HasAnnotationWithValue(testObject, "a.missing/annotation", "")
	if checkMissingAnnotation != false {
		t.Fatalf("HasAnnotationWithValue() found a non-existent annotation")
	}
}

func TestHasLabel(t *testing.T) {
	checkExistingLabel := HasLabel(testObject, "a.test/label")
	if checkExistingLabel != true {
		t.Fatalf("HasLabel() failed to find existing label")
	}

	checkMissingLabel := HasLabel(testObject, "bar")
	if checkMissingLabel != false {
		t.Fatalf("HasLabel() found a non-existent label")
	}
}

func TestHasLabelWithValue(t *testing.T) {
	checkLabelProperValue := HasLabelWithValue(testObject, "a.test/label", "foo")
	if checkLabelProperValue != true {
		t.Fatalf("HasLabelWithValue() failed to find value of existing label")
	}

	checkLabelWrongValue := HasLabelWithValue(testObject, "a.test/label", "bar")
	if checkLabelWrongValue != false {
		t.Fatalf("HasLabelWithValue() found the wrong value of an existing label")
	}

	checkMissingLabel := HasLabelWithValue(testObject, "a.missing/label", "")
	if checkMissingLabel != false {
		t.Fatalf("HasLabelWithValue() found a non-existent label")
	}
}
