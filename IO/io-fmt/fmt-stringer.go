package main

import(
	"fmt"
	"bytes"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Sex int
}

func (this *Person) String() string {
	// NewBufferString使用字符串创建一个新的缓存区
	buffer := bytes.NewBufferString("This is ")
	// WriteString 将字符串住追加到缓存区。
	buffer.WriteString(this.Name + ",")
	if this.Sex == 0 {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}
	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(this.Age))
	buffer.WriteString(" years old.")
	// String以字符串形式返回缓冲区未读部分的内容。如果Buffer是nil指针，则返回“<nil>”。
	return buffer.String()
}

func (this *Person) Format(f fmt.State, c rune) {
	if c == 'L' {
		f.Write([]byte(this.String()))
		f.Write([]byte(" Person has three fields."))
	} else {
		f.Write([]byte(fmt.Sprintln(this.String())))
	}
}
func main() {
	p := &Person{"luyuan", 11, 0}
	fmt.Println(p)
	fmt.Printf("%L", p)
}