# 后端项目结构说明
├─LICENSE
├─README.md
├─backend
│  ├─.env   
│  ├─README.md
│  ├─api    # 统一的接口契约层(Protobuf/Buf代码生成)
│  │  ├─README.md  
│  │  └─proto   
│  │     ├─ai
│  │     │  └─v1
│  │     ├─collab
│  │     │  └─v1
│  │     ├─common
│  │     │  └─v1
│  │     ├─doc
│  │     │  └─v1
│  │     ├─iam
│  │     │  └─v1
│  │     └─notify
│  │        └─v1
│  ├─app    # backend/app/{domain}/{process} 
│  │  ├─ai  # domain：业务域(iam/doc/collab/ai/notify)
│  │  │  └─service  # process：进程类型(service对外提供API；job跑后台任务)
│  │  │     ├─cmd
│  │  │     ├─configs
│  │  │     └─internal
│  │  ├─collab
│  │  │  └─service
│  │  │     ├─cmd
│  │  │     ├─configs
│  │  │     └─internal
│  │  ├─doc
│  │  │  ├─job
│  │  │  │  └─cmd
│  │  │  └─service
│  │  ├─gateway
│  │  │  └─service
│  │  ├─iam
│  │  │  └─service
│  │  └─notify
│  │     └─service
│  ├─deployments    # 部署和可观测
│  ├─pkg    # 跨服务复用的公共库
│  │  ├─constants
│  │  ├─jwt
│  │  └─middleware
│  └─script # 环境初始化/一键部署脚本

Atlas 后端采用 Monorepo 管理多微服务，按“领域 + 进程”划分目录，接口契约统一放在 api，服务实现放在 backend/app/**，公共能力在 pkg。