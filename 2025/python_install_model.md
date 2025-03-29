AI: Unsloth + Llama 3 微调实践,基于Colab

## **引言**

Meta AI 的 Llama 3 开源大语言模型发布后，引起了广泛关注。然而，直接使用 Llama 3 完整版（如 8B 参数）对硬件，尤其是 GPU 有较高要求。别担心，这里将介绍如何利用 Unsloth 库，在资源有限的条件下（例如 Google Colab 的免费 Tesla T4 GPU），高效微调 Llama 3 的较小版本（1B 和 3B 参数），让它们也能胜任对话任务。本文提供完整、可运行的代码，并详细讲解每一步，确保不熟悉的同学也能上手。

## **什么是大语言模型（LLM）？**

在深入之前，我们先简单了解一下什么是大语言模型。我们可以把它们想象成超级聪明的“鹦鹉”，它们通过阅读大量的文本（比如书籍、文章、网页）来学习语言。学完之后，它们就能做很多事情，比如：

*   **回答问题：** 就像一个知识渊博的助手。
*   **写文章：** 可以写各种类型的文本，比如新闻、故事、代码。
*   **翻译语言：** 把一种语言翻译成另一种语言。
*   **进行对话：** 就像一个聊天机器人，可以和我们对话。

Llama 3 就是这样一种大语言模型，而且它是开源的，这意味着我们可以免费使用和修改它。

## **为什么需要“微调”？**

直接使用像 Llama 3 这样的大模型，虽然它很强大，但可能在特定任务上表现不够完美。比如，我们希望它专门用于客服对话，或者用于写特定风格的诗歌。这时，我们就需要“微调”。

微调就像是对大模型进行“再教育”。我们给它提供一些特定领域的例子（比如客服对话记录，或者某个诗人的诗歌作品），让它在这个领域变得更专业。微调之后，模型在特定任务上的表现会更好。

## **Unsloth：加速训练的秘密武器**

微调大模型通常需要大量的计算资源（特别是 GPU）和时间。Unsloth 就是一个专门为解决这个问题而设计的工具。它就像一个“加速器”，可以让训练过程变得更快、更省资源。

Unsloth 的主要优点：

*   **更快：** 训练速度比常用的 Hugging Face Transformers 库快很多（有时甚至快 30 倍！）。
*   **更省内存：** 训练所需的 GPU 内存大大减少。
*   **更易用：** 使用起来比较简单，代码更简洁。

Unsloth 之所以这么厉害，是因为它使用了各种优化技术，比如：

*   **Flash Attention-2：** 一种更高效的注意力机制，是 Transformer 模型的核心。
*   **量化：** 把模型的参数（可以理解为模型的“知识”）用更小的格式存储，从而减少内存占用和计算量。

## **LoRA：轻量级微调的利器**

LoRA (Low-Rank Adaptation) 是一种非常流行的微调方法。它最大的特点是“轻量级”。

传统的微调方法需要调整模型的所有参数，而 LoRA 只调整其中一小部分。这就好比，我们学习英语不需要同时学习驾驶，只需要做学习英语相关的活动，以调整相关的大脑神经连接，就能让我们的大脑逐渐熟悉英语。

LoRA 的优点：

*   **高效：** 训练速度快，需要的资源少。
*   **性能好：** 在很多情况下，LoRA 的效果可以媲美全参数微调。

## **动手实践：用 Unsloth 微调 Llama 3**

现在，我们将一步步展示如何在 Colab 中使用 Unsloth 微调 Llama 3。

1.  **环境准备：**

    *   **重要提示：**  Colab 环境有时会出现 PyTorch 相关的问题。为了避免这些问题，我们将使用一种特殊的安装方式，而不是直接 `pip install unsloth`。
    *   以下代码将安装必要的库，包括 `bitsandbytes`、`accelerate`、`xformers`、`peft`、`trl`、`triton`、`sentencepiece`、`protobuf`、`datasets`、`huggingface_hub`、`hf_transfer` 以及 `unsloth`。
    *   **特别注意：** 这里的安装命令是经过实践验证、可在 Colab 中稳定运行的版本。

    ```python
    %%capture
    # Colab 中特殊的安装方式，避免 PyTorch 相关问题
    !pip install --no-deps bitsandbytes accelerate xformers==0.0.29 peft "trl<0.15.0" triton
    !pip install --no-deps cut_cross_entropy unsloth_zoo
    !pip install sentencepiece protobuf datasets huggingface_hub hf_transfer
    !pip install --no-deps unsloth
    ```

2.  **加载模型和分词器：**

    *   使用 Unsloth 的 `FastLanguageModel.from_pretrained` 函数加载 Llama 3 模型和分词器。
    *   `model_name`：指定要加载的模型。这里我们选择 `"unsloth/Llama-3.2-3B-Instruct"`，也可以尝试 `"unsloth/Llama-3.2-1B-Instruct"`。
    *  `max_seq_length`：设置模型能处理的最大文本长度（这里是 2048）。
    *   `dtype`：设置模型参数的数据类型。`None` 表示自动检测（Colab T4 GPU 通常会使用 `float16`）。
    *   `load_in_4bit`：启用 4 位量化，减少显存占用。

    ```python
    from unsloth import FastLanguageModel
    import torch

    max_seq_length = 2048  # 最大序列长度
    dtype = None  # 自动检测数据类型
    load_in_4bit = True  # 启用 4 位量化

    model, tokenizer = FastLanguageModel.from_pretrained(
        model_name="unsloth/Llama-3.2-3B-Instruct",  # 或 "unsloth/Llama-3.2-1B-Instruct"
        max_seq_length=max_seq_length,
        dtype=dtype,
        load_in_4bit=load_in_4bit,
        # token = "hf_...",  # 如果使用 Meta 官方的 Llama 2 模型，需要 Hugging Face token
    )
    ```

3.  **添加 LoRA 适配器：**

    *   使用 `FastLanguageModel.get_peft_model` 函数添加 LoRA 适配器。
    *   `r`：LoRA 的秩（rank）。通常选择 8、16、32、64 或 128。
    *   `target_modules`：指定要应用 LoRA 的模块。
    *   `lora_alpha` 和 `lora_dropout`：LoRA 的超参数。
    *   `use_gradient_checkpointing`：使用梯度检查点，进一步减少显存占用。`"unsloth"` 是 Unsloth 提供的优化版本。

    ```python
    model = FastLanguageModel.get_peft_model(
        model,
        r=16,  # LoRA rank
        target_modules=[
            "q_proj",
            "k_proj",
            "v_proj",
            "o_proj",
            "gate_proj",
            "up_proj",
            "down_proj",
        ],
        lora_alpha=16,
        lora_dropout=0,  # 支持任意值，但 0 已优化
        bias="none",  # 支持任意值，但 "none" 已优化
        use_gradient_checkpointing="unsloth",  # "unsloth" 使用更少显存
        random_state=3407,
        use_rslora=False,  # 支持 rank-stabilized LoRA
        loftq_config=None,  # 支持 LoftQ
    )
    ```

4.  **数据准备：**

    *   **Llama 3.1 对话模板：**  Llama 3 使用特定的对话模板来组织多轮对话。
    *   使用 `get_chat_template` 函数获取正确的聊天模板。
      ```
      <|begin_of_text|><|start_header_id|>user<|end_header_id|>

      Hello!<|eot_id|><|start_header_id|>assistant<|end_header_id|>

      Hey there! How are you?<|eot_id|><|start_header_id|>user<|end_header_id|>

      I'm great thanks!<|eot_id|>
      ```

    *   **数据集：**  使用 `mlabonne/FineTome-100k` 数据集，这是一个 ShareGPT 风格的数据集。
    *   **数据格式转换：**  使用 `standardize_sharegpt` 函数将 ShareGPT 格式转换为 Hugging Face 的通用格式：
        *   ShareGPT 格式：`{"from": "...", "value": "..."}`
        *   Hugging Face 格式：`{"role": "...", "content": "..."}`
    *   **应用聊天模板：** 使用 `formatting_prompts_func` 函数将聊天模板应用到数据集。

    ```python
    from unsloth.chat_templates import get_chat_template

    tokenizer = get_chat_template(
        tokenizer,
        chat_template="llama-3.1",
    )
    def formatting_prompts_func(examples):
        convos = examples["conversations"]
        texts = [
            tokenizer.apply_chat_template(
                convo, tokenize=False, add_generation_prompt=False
            )
            for convo in convos
        ]
        return {"text": texts}
    
    from datasets import load_dataset
    
    dataset = load_dataset("mlabonne/FineTome-100k", split="train")

    from unsloth.chat_templates import standardize_sharegpt
    
    dataset = standardize_sharegpt(dataset)
    dataset = dataset.map(formatting_prompts_func, batched=True)

    ```

5.  **查看数据示例：**

    *   查看数据集中的对话示例，以及应用聊天模板后的文本。

    ```python
    #原始对话
    print(dataset[5]["conversations"])
    #应用模版后的对话
    print(dataset[5]["text"])
    ```

6.  **训练模型：**

    *   使用 Hugging Face TRL 库的 `SFTTrainer` 进行训练。
    *   `DataCollatorForSeq2Seq`：用于将数据整理成批次。
    *   `TrainingArguments`：配置训练参数，如批次大小、梯度累积、学习率、优化器等。
        *   `per_device_train_batch_size`: 每个 GPU 的批次大小
        *   `gradient_accumulation_steps`: 梯度累积步数
        *   `warmup_steps`: 预热步数
        *   `max_steps`: 最大训练步数
        *   `learning_rate`: 学习率
        *   `fp16 / bf16`: 是否使用 FP16 或 BF16 混合精度训练
        *   `logging_steps`: 日志记录步数
        *   `optim`: 优化器
        *   `weight_decay`: 权重衰减
        *   `lr_scheduler_type`: 学习率调度器类型
        *   `seed`: 随机种子
        *  `output_dir`: 输出目录
    *  `train_on_responses_only`: 仅训练助理的回复，忽略用户提问的损失

    ```python
    from trl import SFTTrainer
    from transformers import TrainingArguments, DataCollatorForSeq2Seq
    from unsloth import is_bfloat16_supported

    trainer = SFTTrainer(
        model=model,
        tokenizer=tokenizer,
        train_dataset=dataset,
        dataset_text_field="text",
        max_seq_length=max_seq_length,
        data_collator=DataCollatorForSeq2Seq(tokenizer=tokenizer),
        dataset_num_proc=2,
        packing=False,  # 对于短序列，可以加快训练速度
        args=TrainingArguments(
            per_device_train_batch_size=2,
            gradient_accumulation_steps=4,
            warmup_steps=5,
            # num_train_epochs=1,  # 完整训练一轮
            max_steps=60,  # 为了演示，只训练 60 步
            learning_rate=2e-4,
            fp16=not is_bfloat16_supported(),
            bf16=is_bfloat16_supported(),
            logging_steps=1,
            optim="adamw_8bit",
            weight_decay=0.01,
            lr_scheduler_type="linear",
            seed=3407,
            output_dir="outputs",
            report_to="none",  # 如果使用 WandB 等，可以在这里配置
        ),
    )

    from unsloth.chat_templates import train_on_responses_only
    trainer = train_on_responses_only(
        trainer,
        instruction_part="<|start_header_id|>user<|end_header_id|>\n\n",
        response_part="<|start_header_id|>assistant<|end_header_id|>\n\n",
    )
    ```

7.  **验证 Masking：**

```python
tokenizer.decode(trainer.train_dataset[5]["input_ids"])
space = tokenizer(" ", add_special_tokens = False).input_ids[0]
tokenizer.decode([space if x == -100 else x for x in trainer.train_dataset[5]["labels"]])
```
8. **查看显存和训练时间统计：**
```python
gpu_stats = torch.cuda.get_device_properties(0)
start_gpu_memory = round(torch.cuda.max_memory_reserved() / 1024 / 1024 / 1024, 3)
max_memory = round(gpu_stats.total_memory / 1024 / 1024 / 1024, 3)
print(f"GPU = {gpu_stats.name}. Max memory = {max_memory} GB.")
print(f"{start_gpu_memory} GB of memory reserved.")

trainer_stats = trainer.train()

used_memory = round(torch.cuda.max_memory_reserved() / 1024 / 1024 / 1024, 3)
used_memory_for_lora = round(used_memory - start_gpu_memory, 3)
used_percentage = round(used_memory / max_memory * 100, 3)
lora_percentage = round(used_memory_for_lora / max_memory * 100, 3)
print(f"{trainer_stats.metrics['train_runtime']} seconds used for training.")
print(
    f"{round(trainer_stats.metrics['train_runtime']/60, 2)} minutes used for training."
)
print(f"Peak reserved memory = {used_memory} GB.")
print(f"Peak reserved memory for training = {used_memory_for_lora} GB.")
print(f"Peak reserved memory % of max memory = {used_percentage} %.")
print(f"Peak reserved memory for training % of max memory = {lora_percentage} %.")

```

9.  **推理（与模型对话）：**

    *   启用 Unsloth 的 2 倍速推理模式。
    *   使用 `apply_chat_template` 方法构建对话提示。
    *   使用 `model.generate` 生成回复。

    ```python
    from unsloth.chat_templates import get_chat_template
    tokenizer = get_chat_template(
      tokenizer,
      chat_template = "llama-3.1",
    )
    FastLanguageModel.for_inference(model)  # 启用 2 倍速推理

    messages = [
        {"role": "user", "content": "Continue the fibonnaci sequence: 1, 1, 2, 3, 5, 8,"},
    ]
    inputs = tokenizer.apply_chat_template(
        messages,
        tokenize=True,
        add_generation_prompt=True,  # 生成时必须添加
        return_tensors="pt",
    ).to("cuda")

    outputs = model.generate(
        input_ids=inputs,
        max_new_tokens=64,
        use_cache=True,
        temperature=1.5,
        min_p=0.1,
    )
    tokenizer.batch_decode(outputs)
    ```
    * 使用`TextStreamer`流式输出
    ```python
    FastLanguageModel.for_inference(model) # Enable native 2x faster inference

    messages = [
        {"role": "user", "content": "Continue the fibonnaci sequence: 1, 1, 2, 3, 5, 8,"},
    ]
    inputs = tokenizer.apply_chat_template(
        messages,
        tokenize = True,
        add_generation_prompt = True, # Must add for generation
        return_tensors = "pt",
    ).to("cuda")

    from transformers import TextStreamer
    text_streamer = TextStreamer(tokenizer, skip_prompt = True)
    _ = model.generate(input_ids = inputs, streamer = text_streamer, max_new_tokens = 128,
                       use_cache = True, temperature = 1.5, min_p = 0.1)
    ```

10. **保存和加载模型：**
    *   保存LoRA模型

    ```python
    model.save_pretrained("lora_model")  # Local saving
    tokenizer.save_pretrained("lora_model")
    # model.push_to_hub("your_name/lora_model", token = "...") # Online saving
    # tokenizer.push_to_hub("your_name/lora_model", token = "...") # Online saving
    ```

     *  加载LoRA模型

    ```python
    if False:
        from unsloth import FastLanguageModel
        model, tokenizer = FastLanguageModel.from_pretrained(
            model_name = "lora_model", # YOUR MODEL YOU USED FOR TRAINING
            max_seq_length = max_seq_length,
            dtype = dtype,
            load_in_4bit = load_in_4bit,
        )
        FastLanguageModel.for_inference(model) # Enable native 2x faster inference

    messages = [
        {"role": "user", "content": "Describe a tall tower in the capital of France."},
    ]
    inputs = tokenizer.apply_chat_template(
        messages,
        tokenize = True,
        add_generation_prompt = True, # Must add for generation
        return_tensors = "pt",
    ).to("cuda")

    from transformers import TextStreamer
    text_streamer = TextStreamer(tokenizer, skip_prompt = True)
    _ = model.generate(input_ids = inputs, streamer = text_streamer, max_new_tokens = 128,
                        use_cache = True, temperature = 1.5, min_p = 0.1)
    ```

## **总结**

本文详细介绍了如何使用 Unsloth 库在 Google Colab 环境中微调 Llama 3 模型，并提供了完整、可运行的代码。通过 Unsloth 的加速和 LoRA 轻量级微调技术，即使在资源有限的情况下，也能高效训练出个性化的对话 AI。我们重点强调了在 Colab 中安装 Unsloth 的正确方法，避免了常见的 PyTorch 兼容性问题。

