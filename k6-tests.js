import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 1000 },
    { duration: '1m', target: 10000 },
  ],
  thresholds: {
    http_req_failed: ['rate<0.01'], // http errors should be less than 1%
    http_req_duration: ['avg<200', 'p(95)<200'], // 95 percent of response times must be below 500ms
  },
};

export function setup() {
  let params = {client_id: "service-client",  client_secret: "8oQmmPpyQqWMAnMTMIuEyQyAZ7m2GKCI", grant_type: "password", username: "yoeri@mail.com", password: "1qazxsw2"}
  const res = http.post('https://lemur-18.cloud-iam.com/auth/realms/anonymizer-iam/protocol/openid-connect/token', params)
  return res.json("access_token")
}

export default function getpage(accesstoken) {
  let params = {headers: { 'Authorization': "Bearer "+accesstoken }}
  const res = http.get('http://34.79.251.22/pages/12', params);
  sleep(10);
}
