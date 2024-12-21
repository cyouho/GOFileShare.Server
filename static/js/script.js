document.addEventListener('DOMContentLoaded', () => {
    // 获取文件列表
    function listFiles() {
        fetch('/api/files')
            .then(response => response.json())
            .then(files => {
                const fileList = document.getElementById('files');
                fileList.innerHTML = ''; // 清空现有文件列表

                files.forEach(file => {
                    const li = document.createElement('li');
                    const a = document.createElement('a');
                    a.href = `/api/download/${encodeURIComponent(file)}`;
                    a.textContent = file;
                    li.appendChild(a);
                    fileList.appendChild(li);
                });
            })
            .catch(error => console.error('Error fetching files:', error));
    }

    // 文件上传表单提交
    document.getElementById('upload').addEventListener('submit', (event) => {
        event.preventDefault();

        const formData = new FormData(event.target);

        fetch('/api/upload', {
            method: 'POST',
            body: formData,
        })
        .then(response => response.json())
        .then(data => {
            if (data.message) {
                alert(data.message);
                listFiles(); // 刷新文件列表
            }
        })
        .catch(error => console.error('Error uploading file:', error));
    });

    // 初始加载时获取文件列表
    listFiles();
});