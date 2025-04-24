#!/bin/bash

PORT=3003
TIMEOUT=30

echo "⏳ Aguardando processo escutando na porta $PORT..."

for i in $(seq 1 $TIMEOUT); do
  if ss -ltn | grep -q ":$PORT"; then
    echo "✅ Porta $PORT está em escuta"
    break
  fi
  echo "Tentativa $i: porta $PORT ainda não está escutando..."
  sleep 1
done

if ! ss -ltn | grep -q ":$PORT"; then
  echo "❌ Nenhum processo escutando na porta $PORT após $TIMEOUT segundos."
  exit 1
fi

chmod +x monitor-mem.sh

./monitor-mem.sh $PORT &
MONITOR_PID=$!

k6 run tests/load-tests.js

kill $MONITOR_PID
echo "✅ Teste finalizado e monitoramento encerrado."
