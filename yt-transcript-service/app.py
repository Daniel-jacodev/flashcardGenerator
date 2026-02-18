from flask import Flask, request, jsonify
from youtube_transcript_api import YouTubeTranscriptApi
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

@app.route('/transcript', methods=['POST'])
def get_transcript():
    data = request.json
    video_url = data.get('url')
    try:
        # Extrai o ID do vídeo
        video_id = video_url.split("v=")[1].split("&")[0] if "v=" in video_url else video_url.split("/")[-1]
        
        transcript = YouTubeTranscriptApi.get_transcript(video_id, languages=['pt', 'en'])
        full_text = " ".join([t['text'] for t in transcript])
        
        return jsonify({"transcript": full_text})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    print("🚀 Microsserviço Python rodando em http://localhost:5000")
    app.run(port=5000)