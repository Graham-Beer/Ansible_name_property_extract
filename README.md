# Ansible_name_property_extract
## Version 0.0.3

A Go command line tool to extract the `- name` property to allow quick overview of steps performed.
Compiled folder 'dist' contains linux version. Feel free to convert to other OS.

### usage
Takes file extensions of `.yml` and `.yaml`. Otherwise will throw error.

#### long version
`./dist/YamlNameParser parse --file somefile.yml`
#### short version
`./dist/YamlNameParser yp -f somefile.yaml`
#### aditional switches
`--listview` is a boolean switch value. If enabled, output will have `- ` infront of text.  
`./dist/YamlNameParser parse --file somefile.yml --listview`
#### short version of switch
`./dist/YamlNameParser yp -f somefile.yaml -l`
#### help
`./dist/YamlNameParser parse --help`
`./dist/YamlNameParser --help`
#### Version
`./dist/YamlNameParser --version`

## Version 0.0.4
Add nested check for `name` property under `tasks`.
