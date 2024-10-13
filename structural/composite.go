package structural

import "fmt"

type IComponent interface{
  GetDetails(string) string
}

type File struct {
  Name string
  Size int
}

func (f *File) GetDetails(indent string) string {
  return fmt.Sprintf("%sFile: %s (Size: %d)\n", indent, f.Name, f.Size)
}

type Folder struct {
  Name string
  components []IComponent
}

func (f *Folder) GetDetails(indent string) string {
 result := fmt.Sprintf("%sFolder: %s\n", indent, f.Name)
 for _,c := range f.components {
    result += c.GetDetails(indent + "    ")
  }
  return result
}

func (f *Folder) Add(c IComponent){
  f.components = append(f.components, c)
}

func (f *Folder) Remove(c IComponent){
  for i, v := range f.components {
    if v == c {
      f.components = append(f.components[:i],f.components[i+1:]...)
    }
  }
}
