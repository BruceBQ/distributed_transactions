const a = [
  {
    level: "info",
    ts: "2022-10-06T10:26:07.526+0700",
    caller: "dtmutil/utils.go:118",
    msg: ' 0ms 200 GET /api/dtmsvr/query?gid=h947tBNJUnWYefBCFbVB86 {"branches":[{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:02.348050768+07:00","gid":"h947tBNJUnWYefBCFbVB86","url":"http://localhost:8080/order/cancel","bin_data":"eyJhbW91bnQiOjEwfQ==","branch_id":"01","op":"compensate","status":"prepared"},{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:02.358482142+07:00","gid":"h947tBNJUnWYefBCFbVB86","url":"http://localhost:8080/order/create","bin_data":"eyJhbW91bnQiOjEwfQ==","branch_id":"01","op":"action","status":"succeed","finish_time":"2022-10-06T10:24:02.358482142+07:00"},{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:02.348050768+07:00","gid":"h947tBNJUnWYefBCFbVB86","bin_data":"bnVsbA==","branch_id":"02","op":"compensate","status":"prepared"},{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:02.348050768+07:00","gid":"h947tBNJUnWYefBCFbVB86","url":"http://localhost:8080/customer/verify","bin_data":"bnVsbA==","branch_id":"02","op":"action","status":"prepared"}],"transaction":{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:58.061498274+07:00","gid":"h947tBNJUnWYefBCFbVB86","trans_type":"saga","steps":[{"action":"http://localhost:8080/order/create","compensate":"http://localhost:8080/order/cancel"},{"action":"http://localhost:8080/customer/verify","compensate":""}],"payloads":["{\\"amount\\":10}","null"],"status":"failed","protocol":"http","rollback_time":"2022-10-06T10:24:58.061498274+07:00","options":"{\\"concurrent\\":false}","next_cron_interval":10,"next_cron_time":"2022-10-06T10:25:07.77399407+07:00","wait_result":true,"concurrent":false}}',
  },
  {
    level: "info",
    ts: "2022-10-06T10:28:22.083+0700",
    caller: "dtmutil/utils.go:118",
    msg: ' 0ms 200 GET /api/dtmsvr/all?current=1&limit=10 {"next_position":"","transactions":[{"ID":0,"create_time":"2022-10-06T10:24:02.348050768+07:00","update_time":"2022-10-06T10:24:58.061498274+07:00","gid":"h947tBNJUnWYefBCFbVB86","trans_type":"saga","steps":[{"action":"http://localhost:8080/order/create","compensate":"http://localhost:8080/order/cancel"},{"action":"http://localhost:8080/customer/verify","compensate":""}],"payloads":["{\\"amount\\":10}","null"],"status":"failed","protocol":"http","rollback_time":"2022-10-06T10:24:58.061498274+07:00","options":"{\\"concurrent\\":false}","next_cron_interval":10,"next_cron_time":"2022-10-06T10:25:07.77399407+07:00","wait_result":true,"concurrent":false}]}',
  },
];
