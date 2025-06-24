import http from 'k6/http';
import { check } from 'k6';
import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

export let options = {
  vus: 50,          // 50 usuários simultâneos
  duration: '5s',   // duração total do teste
};

export default function () {
  const apelido = `apelido-${uuidv4()}`;
  const nome = `Pessoa ${uuidv4()}`;
  const nascimento = '1990-01-01';  // data fixa, pode variar se quiser
  const stack = ['Go', 'Node.js'];

  const payload = JSON.stringify({
    apelido,
    nome,
    nascimento,
    stack,
  });

  const headers = { 'Content-Type': 'application/json' };

  const res = http.post('http://localhost:9999/programadores', payload, { headers });

  const success = check(res, {
    'status é 201': (r) => r.status === 201,
  })

  if (!success) {
    console.log(`❌ Erro na requisição: status ${res.status}, body: ${res.body}`)
  }
}
