package main

import (
	"testing"
)

func TestApp_GetPreview(t *testing.T) {
	app := NewApp()
	app.setFiles("/Users/qiaoleilei/Downloads/images")
	app.setPreview()
	app.SetRules(&Rules{
		RenameType: AddNumber,
		NumberObj: &NumberObj{
			NewName: "my_image",
			Suffix:  underLine,
		},
	})
	app.ExecsRename()
}
