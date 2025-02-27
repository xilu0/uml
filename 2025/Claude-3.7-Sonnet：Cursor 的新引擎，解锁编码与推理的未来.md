Claude-3.7-Sonnet：Cursor 的新引擎，解锁编码与推理的未来

## 引言
claude-3.7-sonnet 是 Anthropic 最新发布的大型语言模型，于 2025 年 2 月 24 日推出，并已集成到 Cursor AI 平台中。Cursor 是一个 AI 驱动的集成开发环境（IDE），旨在通过 AI 增强开发者的生产力。claude-3.7-sonnet 被誉为 Anthropic 最智能的模型，引入了混合推理能力，支持快速和逐步推理。接下来将探讨该模型的参数、与前代模型的进展以及对 Cursor 用户的影响。

## 模型参数与 Cursor 集成
在使用 claude-3.7-sonnet 时，开发者需要在 Cursor 中选择该模型，并可能调整某些设置。以下是相关参数的概述：

- **模型选择**：在 Cursor 的设置中选择 claude-3.7-sonnet，确保使用最新版本。
- **思维模式**：用户可以选择标准模式（快速回答）或扩展思维模式（逐步推理），后者特别适合复杂问题。
- **输出长度**：该模型支持高达 128,000 个输出标记，显著高于前代模型，适合生成详细的代码和规划。
- **API 密钥管理**：需要从 Anthropic 获取 API 密钥并在 Cursor 中配置，以访问该模型。

然而，具体参数如温度（影响创造性）、最大标记数或停止序列可能由 Cursor 内部处理，用户无法直接调整。这一点从 Cursor 的官方文档和用户指南中未明确提及，但可以推测为平台优化的一部分。

以下是与 claude-3.7-sonnet 相关的一些技术参数，基于 Amazon Bedrock 的文档：

| **参数**               | **详情**                                                                                     | **具体数值/值**                     |
|-------------------------|----------------------------------------------------------------------------------------------|--------------------------------------|
| 输出标记               | 每请求支持高达 128,000 个输出标记，64,000 个一般可用，64,000-128,000 为 beta 阶段           | 128,000, 64,000                     |
| 最大标记               | 提示标记 + max_tokens 不得超过上下文窗口，否则报错                                           | -                                    |
| 思考预算标记           | 最小 1,024 标记，Anthropic 建议至少 4,000 以确保全面推理                                     | 1,024, 4,000                         |
| 扩展输出长度 beta      | 通过 anthropic-beta 参数启用                                                                 | -                                    |
| 计算机使用 beta 工具   | 新增动作：滚动、等待、左键按下、左键松开、按住键、三击                                       | -                                    |
| 请求有效载荷最大大小   | InvokeModel/InvokeModelWithResponseStream 的最大大小                                         | 20MB                                 |
| 图像支持               | 最大 3.75MB，最大高度/宽度 8,000 像素，支持格式：image/jpeg, image/png, image/webp, image/gif | 3.75MB, 8,000 像素                   |

尽管这些参数在 Cursor 中可能不直接可调，但它们反映了模型的潜在能力，开发者可以通过提示优化输出。

## 与前代模型的进展
claude-3.7-sonnet 相较于前代模型（如 claude-3.5-sonnet）在多个方面取得了显著进展：

- **编码能力**：在 SWE-bench Verified 基准测试中，标准模式得分达 70.3%，领先行业，特别适合处理真实世界的软件工程任务。
- **推理能力**：作为首个混合推理模型，它支持快速和逐步推理，用户可以根据需要选择模式。
- **输出容量**：支持高达 128,000 个输出标记，是 claude-3.5-sonnet 的 15 倍以上，适合生成详细的代码和规划。
- **拒绝率降低**：不必要拒绝减少了 45%，使模型更易于合作，减少了因伦理边界误判而拒绝回答的情况。

这些进展基于 Anthropic 的官方公告和基准测试结果，例如 [Anthropic's Announcement of claude-3.7-sonnet]()(https://www.anthropic.com/news/claude-3-7-sonnet) 和 [Amazon Bedrock Integration]()(https://aws.amazon.com/blogs/aws/anthropics-claude-3-7-sonnet-the-first-hybrid-reasoning-model-is-now-available-in-amazon-bedrock/)。

此外，该模型在 TAU-bench（测试 AI 代理在复杂任务中的表现）和其他基准测试中也表现出色，尤其是在数学和科学领域，扩展思维模式显著提升了性能。

## 对 Cursor 用户的影响
claude-3.7-sonnet 的集成对 Cursor 用户带来了多方面的影响，增强了开发体验：

- **更准确的代码建议**：模型在编码任务上的优异表现意味着开发者可以获得更准确的代码补全和建议，减少手动调整的时间。
- **深入的代码分析**：扩展思维模式允许模型逐步推理，特别适合调试复杂代码或理解大型代码库。
- **高效的任务完成**：支持快速和详细两种模式，开发者可以根据任务需求灵活切换，提高工作效率。
- **自动化潜力**：模型支持实验性的计算机使用功能（如移动光标、点击按钮），虽然仍在 beta 阶段，但未来可能进一步自动化开发任务，如运行终端命令或编辑文件。

这些功能基于 Cursor 的官方特性页面 [Cursor AI Features]()(https://www.cursor.com/en/features) 和用户指南。

一个意想不到的细节是，该模型还支持图像输入，最大 3.75MB，适合视觉相关任务，如分析建筑图纸，这可能为 Cursor 用户在多模态开发中提供额外支持。

## 结论
claude-3.7-sonnet 的引入显著提升了 Cursor 的功能，使其成为开发者更强大的工具。其混合推理能力、长输出支持和优异的编码表现为开发流程带来了效率和深度的提升。尽管某些参数可能不可直接调整，但通过提示优化和模式选择，开发者可以充分利用该模型的能力。

## 关键引用
- [Anthropic's Announcement of claude-3.7-sonnet]()(https://www.anthropic.com/news/claude-3-7-sonnet)
- [Cursor AI Features]()(https://www.cursor.com/en/features)
- [Amazon Bedrock Integration]()(https://aws.amazon.com/blogs/aws/anthropics-claude-3-7-sonnet-the-first-hybrid-reasoning-model-is-now-available-in-amazon-bedrock/)
- [Step-by-Step Guide to Using Cursor AI with Claude Models]()(https://medium.com/towards-agi/how-to-use-cursor-ai-with-claude-3-5-sonnet-step-by-step-guide-4e1bbdd7bd65)