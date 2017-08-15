# rmbom

Small cli utility to remove the Byte Order Mark (BOM) from files.

## Usage

```sh
# overwrite input file without the BOM
$ rmbom -input bom.json

# specify the output file
$ rmbom -input bom.json -output nobom.json
```
