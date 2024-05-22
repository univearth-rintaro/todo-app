-- テーブル作成
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL,
    done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 初期データ挿入
INSERT INTO todos (title, description, status) VALUES
('買い物', '牛乳、パン、チーズ、卵', 'pending'),
('プロジェクト完了', 'プロジェクトの残りのタスクを完了する', 'in-progress'),
('本を読む', '新しい小説を読み始める', 'completed');