# Project Overview
This project aims to compare the performance differences between Go and R in estimating the standard error of the median using the bootstrap method. We use Go to generate 400 random samples from a normal distribution and save them as a CSV file for R to access. Then, we perform 1,000 bootstrap resamples with replacement in both Go and R, calculating the standard error of the median. At the same time, we record the total execution time and memory usage for both languages to evaluate their efficiency differences.

---

# Program Features
The core features of this project include:
- Generate 400 normally distributed sample data using Go and save them as a CSV file
- Then perform 1,000 bootstrap resamples in both Go and R to calculate the standard error of the median
- Record and output the total runtime and memory usage for each execution

---

# How to Run
Use the following command to run the program: 
- go run bootstrap.go

If everything works correctly, the terminal will display: 
- Saved original sample to original_data.csv
- Estimated SE of Median (Go)
- Total runtime
- Total memory used

Output generated: 
- original_data.csv

---

# Performance Comparison
In this experiment, we performed 1,000 bootstrap resamples in both Go and R using the same original sample. The results are as follows:

Go
- Estimated SE of Median: 0.69662
- Total runtime: 30 ms
- Total memory used: 3226480 bytes

R
- Estimated SE of Median: 0.64843 
- Total runtime: 139.4 ms
- Total memory used: 153392 bytes

In terms of execution time, Go is significantly faster—more than four times faster than R. However, when it comes to memory usage, R is more efficient, using only about 5% of the memory that Go consumes. This difference is mainly due to Go's focus on concurrency and computational speed.

---

# Feasibility Assessment of Replacing R with Go
The experiment on estimating the standard error of the median via bootstrap using identical data revealed that Go runs approximately 4.6 times faster than R. Although Go consumes more memory, its efficient execution significantly reduces computation time in cloud environments, thereby lowering overall costs. For long-running or large-scale concurrent tasks, it is recommended that this consultancy consider adopting Go to help reduce cloud computing expenses.

---

# Use of AI Assistants
- Searched for how to generate original samples using Go and save them as a CSV file
- Searched for how to record execution time and memory usage in R

