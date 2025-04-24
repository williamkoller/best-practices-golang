import http from 'k6/http';
import { check, sleep } from 'k6';
import { Trend, Rate, Counter } from 'k6/metrics';

export const options = {
    stages: [
        { duration: '1m', target: 50 },
        { duration: '3m', target: 100 },
        { duration: '30s', target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(95)<500'],
        http_req_failed: ['rate<0.01'],
    },
};

// Métricas customizadas
const durationTrend = new Trend('custom_duration');
const successRate = new Rate('success_rate');
const failedRequests = new Counter('failed_requests');

const tokens = [
    'c6BX1nk5Y6s:APA91bEb_D8KuCGC_ZTI2pvWIJJZdrymGwrwlWHHA7U6ieKwAK8bV7zpXKSvSDOSgTJn1GnoD49ivjoXoPiLg_P2RuEJfeG2O4PzS_ibR4Xoj0o2bGf-2UHQ7CBdGs3621m8450BdDs0',
    'cwYKyetoSsM:APA91bHXXAkBfBBwC0d261-orrRi_h3vbs1cfzd8GUF8uXZTEKsKbDDx4ygq2tlpBTnareXA63yf225EEW-LEDEhVFNt12Uyupov_oxhwhQEBhD6qfEJqE1srjHrChiV7xkoz-cmCP1y',
    'c32g__9uzFM:APA91bGy-BimZfwlVoKMolN8cBPDYy_gznvxbeb-oiz7sQWcKLtDnqypAJpdj0vE_L8iq9iEywsVue37IOkKkhmK1dNP9aWmAjwFMqu2wAWcRLhoKCiXtwgNZVH5ajr49-S9OfuE0EQX',
    'dTqoSihLFHA:APA91bEoF6eergEyFdcqdPD6dCw1GoopD6LqrjXcLoM18uOhtckYmGGS7I6jPvPQ7AGN5IBSaztJgVUtl3KYsR5vFcQKxQYuPe2hq6-OvM6NN9B_TAAeX0JoIlRfFEOL9Ro7JDiEAb4q',
    'cWL1GYSlb6A:APA91bFDJm0KnzHDpCc53JeY4Rj5kQu6oHDIwMkd0iUG80TqsWSzRpmvJZcioh4XmLandiwobAQSWWPx34eG6_6EljifrPlgYhgy9q2FwH8mTcM8AwT1pC7S_q2XH3NaK82g-atgLJez'
];

export default function () {
    const url = 'http://localhost:3003/tokens';

    // Distribui tokens pela VU e iteração
    const tokenIndex = (__VU - 1) * 100 + __ITER;
    const token = tokens[tokenIndex % tokens.length];

    const payload = JSON.stringify({ token });
    const params = { headers: { 'Content-Type': 'application/json' } };

    const res = http.post(url, payload, params);

    check(res, {
        'status is 201': (r) => r.status === 201,
        'response time < 500ms': (r) => r.timings.duration < 500,
    });

    durationTrend.add(res.timings.duration);
    successRate.add(res.status === 201);
    if (res.status !== 201) failedRequests.add(1);

    sleep(1);
}
