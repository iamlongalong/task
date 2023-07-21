<div align="center">
  <a href="https://taskfile.dev">
    <img src="docs/static/img/logo.svg" width="200px" height="200px" />
  </a>

  <h1>Task</h1>

  <p>
    Task is a task runner / build tool that aims to be simpler and easier to use than, for example, <a href="https://www.gnu.org/software/make/">GNU Make<a>.
  </p>

  <p>
    <a href="https://taskfile.dev/installation/">Installation</a> | <a href="https://taskfile.dev/usage/">Documentation</a> | <a href="https://twitter.com/taskfiledev">Twitter</a> | <a href="https://fosstodon.org/@task">Mastodon</a> | <a href="https://discord.gg/6TY36E39UK">Discord</a>
  </p>
</div>

## 说明

这个项目是从 [go-task/task](https://github.com/go-task/task) 中 fork 过来的。

挺喜欢这个项目的设计，用来做 scripts 管理很方便，目前在工作中用来做一些常用的环境准备，比如安装一套 k8s 环境、安装一些常用工具之类的。

相比于 ansible，这个项目最吸引我的点有两个： 

- 依赖很少，一个 3MB 的二进制就能跑！
- 基于 yaml 的配置，内容很简洁 (没有 ansible 封装的那种自定义的 schema 结构)

但在使用中也遇到了一些问题，主要还是在 环境变量的传递问题！！！

作者在 Q&A 中也解释了不同 shell 间不能共享环境变量，这在同一层级的任务中其实很好理解。

但我的场景是，在一个 task 中通过调用另一个 task，写法如下：

```yaml

tasks:
  parent_task:
    env:
      MY_ENV: "value from parent"
    cmds:
      - echo "this is parent, and MY_ENV is $MY_ENV"
      - task: child_task
  child_task:
    cmds:
      - 'echo "this is child, and I DO NOT HAS MY_ENV !!!  see: $MY_ENV"'
```

这非常令人疑惑 😵😵😵 为什么在子 cmd 中的 task 却不能用父 task 的环境变量呢？

所以，我 fork 了一版，增加了一个 feature: 在子 cmd 中，若使用了 task, 则会继承父 task 的环境变量。

按照原有逻辑，子 task 的环境变量继承顺序，因此修改后的环境变量继承逻辑为:

taskfile envs => dotenv envs => parent task envs => task envs。

但是需要注意，我的改动仅实现了在一个 task 的 cmd 中调用了另一个 task 的情况，**不包含** 使用 dep 的方式。


## TODDO

考虑到实际使用时的场景，后续决定再加一些功能：

- 环境变量支持 default 语法，类似于 `MY_VAR=${MY_VAR:-"var_default"}` 这种，用于子 task 中定义默认的环境变量。 (类似于 vars: MY_VAR: sh: "xxx" 的模式)
- required 环境变量，(类似于 required vars) 

这些改动都是针对环境变量的，我认为一个脚本管理工具，最好是保持兼容性，环境变量 相对于 模板变量 而言，前者的兼容性就更好，大家理解起来也更加符合原来的 shell 语法。


还在纠结要不要加的功能：

- task 的包管理，能够实现依赖收集，类似于从一个 repo 中拉取所需要的脚本文件等

> 我的场景是，启动一台机器然后执行一些任务，因为新启动的机器没有任何 Taskfile 文件及所调用的文件，所以需要下载，那么在哪里下载，下载哪些东西呢？
> 这个功能我现在是用一个 bash script 解决的，行为模式上有点类似于 npx、pipx 这类工具，所以在想要不要直接加到 task 内置 task 中，例如 `task -i pull repo/taskgroup` 或者 `task -i run repo/taskgroup`
> 这样的话，将来集中式的 Taskfile 管理就比较方便了

- 增加 hooks

> 有一些场景中，在执行一个任务时，需要做一些前置的处理，有些需要有后置的处理，比如 打印开始和结束的log、对任务做统计、开始和结束发webhook通知等等。
> hooks 在这类场景中很有用，更好地实现一些公共能力的可复用性。

上面这两项能力对项目的改动都比较大，所以我也拿不定主意，收集下大家的建议吧

如果，你认为上面的功能对你有价值，请在 issue 中留言投个票，并描述一下你的场景~~
