# ✅ 基础功能需求
1. Visit(url string)
插入新历史记录节点。

清空当前节点之后的所有 forward 历史。

2. Back(steps int) string
向前（prev）移动 steps 步。

返回当前页面的 URL。

3. Forward(steps int) string
向后（next）移动 steps 步。

返回当前页面的 URL。

4. Current() string
返回当前页面 URL。

5. PrintHistory()
从头开始打印所有历史。

用 * 标记当前页面。

# 🔥 进阶功能
- 加入访问时间戳（用 time.Time）
- 支持“搜索历史”（关键词匹配 URL）
- 限制历史最大长度（如只保留最近 100 条）
- 实现 Undo() / Redo() 支持