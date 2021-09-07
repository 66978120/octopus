# Octopus Platform

<img src="./logo.png" width="100">

---

[EN](./readme_en.md)

**Octopus**是一款面向多计算场景的一站式融合计算平台。平台主要针对AI、HPC等场景的计算与资源管理的需求来设计，向算力使用用户提供了对数据、算法、镜像、模型与算力等资源的管理与使用功能，方便用户一站式构建计算环境，实现计算。同时，向集群管理人员提供了集群资源管理与监控，计算任务管理与监控等功能，方便集群管理人员对整体系统进行操作与分析。

**Octopus**平台底层基于容器编排平台[Kubernetes](https://kubernetes.io/zh/docs/concepts/overview/what-is-kubernetes) ，充分利用容器敏捷、轻量、隔离等特点来实现计算场景多样性的需求。

## 文档

详细文档请参考[这里](https://octopus.openi.org.cn/docs/introduction/intro)。

## 特点与场景

Octopus具有如下特点：

- **一站式开发**，为用户提供一站式AI、HPC计算场景的开发功能，通过数据管理、模型开发和模型训练，打通计算全链路；
- **方便管理**，为平台管理者提供一站式的资源管理平台，通过资源配置、监控、权限管控等可视化工具，大大降低平台管理者的管理成本；
- **易于部署**，Octopus 支持[Helm](https://helm.sh)方式的快速部署，简化复杂的部署流程；
- **性能优越**，提供高性能的分布式计算体验，通过多方面优化来保证各个环境的流畅运行，同时通过资源调度优化与分布式计算优化，进一步提高模型训练效率；
- **兼容性好**，平台支持异构硬件，如 GPU、NPU、FPGA 等，满足各种不同的硬件集群部署需求，通过支持多种深度学习框架，如 TensorFlow、Pytorch、PaddlePaddle 等，并可以通过自定义镜像方式支持新增框架。

Octopus适合在如下场景中使用：

- 构建大规模 AI 计算平台；
- 希望共享计算资源；
- 希望在统一的环境下完成模型训练；
- 希望使用集成的插件辅助模型训练，提升效率。

## 开始

**Octopus**管理计算资源并针对AI、HPC等场景的计算任务进行优化。通过镜像与容器技术([Docker](https://docs.docker.com))实现计算硬件与软件解耦，从而轻松切换不同计算环境中。

由于Octopus的使用用户通常有两种不同的角色：

- **集群管理员**是计算资源的所有者和维护者。管理员负责集群的部署和可用性。
- **集群用户**是集群计算资源的消费者。根据部署场景，集群用户可以是机器学习和深度学习的研究人员、数据科学家、实验室教师、学生等。

Octopus 为集群用户和管理员提供端到端的手册。

### 对于集群管理员

与集群管理员相关的文档包括如下：

- ***集群部署指南***: 此部分主要提供的内容包括：集群依赖环境与组件的准备与安装、Octopus系统部署指南以及后续系统的升级说明等，以方便安装维护。详细内容请参考[这里](https://octopus.openi.org.cn/docs/deployment/environment) 。

- ***集群管理手册***: 此部分主要介绍集群管理员通过管理系统页面入口进入Octopus管理系统后可进行的操作，主要功能说明包括：平台监控、资源管理、用户管理、机时管理、数据管理、算法管理以及开发与训练管理等功能。详细内容请参考[这里](https://octopus.openi.org.cn/docs/management/intro) 。

### 对于集群用户

与集群用户相关的文档主要如下：

- ***用户使用手册***： 此部分主要介绍集群用户通过Octopus系统页面入口进入Octopus系统后可进行的操作，主要功能说明包括：数据管理、算法管理、镜像管理以及开发与训练管理等功能。详细内容请参考[这里](https://octopus.openi.org.cn/docs/manual/intro) 。

## 如何贡献

详细贡献指南请参考[这里](https://octopus.openi.org.cn/docs/community/contribution) 。

## License

[Apache License](https://octopus.openi.org.cn/docs/community/LICENSE)