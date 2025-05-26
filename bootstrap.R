# 加载内存监控库
library(pryr)

# 设置bootstrap次数
B <- 1000

# 记录开始时间和内存
start_time <- Sys.time()
mem_before <- mem_used()

# 读取Go生成的数据
data <- read.csv("original_data.csv", header = FALSE)
thisSample <- data$V1
n <- length(thisSample)

# 执行bootstrap
bootstrapMedians <- numeric(B)
set.seed(9999)  # 保证结果可复现

for (b in 1:B) {
  resample <- sample(thisSample, n, replace = TRUE)
  bootstrapMedians[b] <- median(resample)
}

# 计算标准误差（标准差）
se_median <- sd(bootstrapMedians)

# 记录结束时间和内存
end_time <- Sys.time()
mem_after <- mem_used()

# 总用时 & 总内存差值
duration_ms <- round(as.numeric(difftime(end_time, start_time, units = "secs")) * 1000, 1)
mem_diff <- mem_after - mem_before

# 打印输出
cat("Estimated SE of Median (R):", round(se_median, 5), "\n")
cat("Total runtime:", duration_ms, "ms\n")
cat("Total memory used:", format(mem_diff, units = "auto"), "\n")
