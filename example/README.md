# End-to-end command line algorithm

Use the `main.go` provided in this folder to run MIDAS algorithm from command line.

The file `main.go` just requires you to specify the `data.csv` file containing containing `src`, `dst` and `timestamps`.

You just have to run:
```bash
go run main.go -input <input-file> -<other-optional-arguments>
``` 

Complete details for the available arguments:
```bash
-alpha float
   Alpha: Temporal Decay Factor. Default is 0.6 (default 0.6)
-buckets int
   Number of buckets. Default is 769 (default 769)
-input string
   Input File. (Required)
-norelations
   To run Midas instead of Midas-R.
-output string
   Output File. Default is scores.txt (default "scores.txt")
-rows int
   Number of rows/hash functions. Default is 2 (default 2)
-undirected
   If graph is undirected.
```

## AUC

To find the AUC for the scores.txt (or output file specified by you using -output argument), 
just run:
```python
python auc.py
```

`auc.py` is also provided in this folder.

## Example usage:

Suppose I have `darpa_processed.csv` file as my data file. Then to run MIDAS-R algorithm I will type:
```bash
go run main.go -input "darpa_processed.csv"  -rows 8 -buckets 512
```

To run MIDAS algorithm I will type:
```bash
go run main.go -input "darpa_processed.csv" -rows 8 -buckets 512 -norelations
```
