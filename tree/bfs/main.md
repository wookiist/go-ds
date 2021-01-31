# BFS (Breadth-First Search)

## 너비 우선 탐색
- 정점들과 같은 레벨에 있는 노드(형제 노드)들을 먼저 탐색하는 방식
- 한 단계씩 내려가면서, 해당 노드와 같은 레벨에 있는 노드(형제 노드)들을 먼저 순회함.

## 그래프를 표현하는 방뻡
- map과 slice를 활용해서 그래프를 표현할 수 있을 것으로 보임.

```go
type Graph map[string][]string
```

## BFS 알고리즘의 구현
- Queue를 사용할 수 있음.
- need_visit 큐와, visited 큐.

1. need_visit 큐에 일단 root node를 넣는다.
2. need_visit 큐에서 하나의 node를 꺼낸다.
3. 꺼낸 node가 visited queue에 존재하는지 확인한다.
    1. 존재하지 않으면, visited queue에 넣는다.
    2. 존재하면, 아무것도 안하고 2번으로 돌아간다.
4. 꺼낸 node에 연결된 node를 need_visit queue에 넣는다.
5. 2로 가서 반복한다.