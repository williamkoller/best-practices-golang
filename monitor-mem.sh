#!/bin/bash

PORT=${1:-3003}
SCRIPT_DIR=$(dirname "$0")
CSV_FILE="$SCRIPT_DIR/mem-usage.csv"

PID=$(lsof -ti tcp:$PORT)

if [ -z "$PID" ]; then
  PID=$(ps aux | grep "api" | grep -v grep | awk '{print $2}' | head -n 1)
fi

if [ -z "$PID" ]; then
  echo "❌ Nenhum processo encontrado escutando na porta $PORT."
  exit 1
fi

echo "✅ Monitorando uso de memória do processo PID=$PID (porta $PORT)"
echo "timestamp,rss_kb" > "$CSV_FILE"

while true; do
  TIMESTAMP=$(date +%s)
  RSS=$(ps -p "$PID" -o rss= | awk '{print $1}')
  if [ -z "$RSS" ]; then
    echo "⚠️  RSS não encontrado para PID $PID, processo pode ter morrido."
    break
  fi
  echo "$TIMESTAMP,$RSS" >> "$CSV_FILE"
  echo "📊 [$TIMESTAMP] Memória usada: ${RSS} KB"
  sleep 1
done
