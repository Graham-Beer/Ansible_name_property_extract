# Ansible_name_property_extract
A Go command line tool to extract the `- name` property to allow quick overview of steps performed.
Compiled folder 'dist' contains linux version. Feel free to convert to other OS.

### usage
Takes file extensions of `.yml` and `.yaml`. Otherwise will throw error.

#### long version
`./dist/YamlNameParser parse --file somefile.yml`
#### short version
`./dist/YamlNameParser yp -f somefile.yaml`
#### help
`./dist/YamlNameParser parse --help`
`./dist/YamlNameParser --help`
#### Version
`./dist/YamlNameParser --version`
