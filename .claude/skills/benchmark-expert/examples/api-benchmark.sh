#!/bin/bash
# Benchmark Expert - Simple API Load Test
# This script is a practical tool for the benchmark-expert skill.

if [ -z "$1" ]; then
  echo "Usage: $0 <url> [requests] [concurrency]"
  exit 1
fi

URL=$1
REQUESTS=${2:-100}
CONCURRENCY=${3:-10}

echo "🚀 Starting benchmark on $URL"
echo "Requests: $REQUESTS | Concurrency: $CONCURRENCY"
echo "---------------------------------------------------"

# Check if ab (Apache Benchmark) is installed
if ! command -v ab &> /dev/null; then
    echo "❌ Error: 'ab' (apache2-utils) is not installed."
    echo "Install it via 'brew install ab' or 'sudo apt install apache2-utils'."
    exit 1
fi

# Run the benchmark and save to a temporary file
TMP_RESULT=$(mktemp)
ab -n $REQUESTS -c $CONCURRENCY $URL > $TMP_RESULT

# Parse key metrics
REQ_PER_SEC=$(grep "Requests per second:" $TMP_RESULT | awk '{print $4}')
TIME_PER_REQ=$(grep "Time per request:" $TMP_RESULT | head -n 1 | awk '{print $4}')
P99_TIME=$(grep " 99%" $TMP_RESULT | awk '{print $2}')
FAILS=$(grep "Failed requests:" $TMP_RESULT | awk '{print $3}')

echo "✅ Benchmark Complete!"
echo "---------------------------------------------------"
echo "Metrics:"
echo "  - Requests per second: $REQ_PER_SEC"
echo "  - Mean latency:        ${TIME_PER_REQ}ms"
echo "  - 99th percentile:     ${P99_TIME}ms"
echo "  - Failed requests:     ${FAILS:-0}"
echo "---------------------------------------------------"

# If LCP/Latency increases by 15% (mocked logic here for baseline check)
# A real implementation would compare with .benchmarks/baseline.json

echo "Generating baseline report..."
mkdir -p .benchmarks
REPORT_FILE=".benchmarks/baseline-$(date +%s).json"
cat <<EOF > $REPORT_FILE
{
  "url": "$URL",
  "requests": $REQUESTS,
  "concurrency": $CONCURRENCY,
  "req_per_sec": $REQ_PER_SEC,
  "p99_latency": $P99_TIME,
  "failed": ${FAILS:-0}
}
EOF

echo "Saved baseline to $REPORT_FILE"
rm $TMP_RESULT
