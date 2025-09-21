import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  vus: 1000,          // 10 virtual users
  duration: "10s",  // run for 10 seconds
};

export default function () {
  const res = http.get("http://localhost:8080/health");

  check(res, {
    "status is 200": (r) => r.status === 200,
    "body is OK": (r) => r.body.includes("OK"),
  });

  sleep(1); 
}
