# DFS (Depth-First Search)

## 깊이 우선 탐색
- 정점의 자식들을 먼저 탐색하는 방식
- 한 노드의 자식을 타고 끝까지 순회한 후, 다시 돌아와서 다른 형제들의 자식을 타고 내려가며 순회함.

## DFS 알고리즘의 구현
- Stack과 queue를 사용
- need_visit 스택(재귀 함수), visited 큐.

1. need_visit 큐에 일단 root node를 넣는다.
2. need_visit 큐에서 하나의 node를 꺼낸다.
3. 꺼낸 node가 visited queue에 존재하는지 확인한다.
    1. 존재하지 않으면, visited queue에 넣는다.
    2. 존재하면, 아무것도 안하고 2번으로 돌아간다.
4. 꺼낸 node에 연결된 node를 need_visit queue에 넣는다.
5. 2로 가서 반복한다.