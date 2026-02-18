from flask import Flask, request, jsonify
from flask_cors import CORS
from youtube_transcript_api import YouTubeTranscriptApi

app = Flask(__name__)
CORS(app)

def extract_video_id(url):
    if "v=" in url:
        return url.split("v=")[1].split("&")[0]
    elif "youtu.be/" in url:
        return url.split("youtu.be/")[1].split("?")[0]
    elif "shorts/" in url:
        return url.split("shorts/")[1].split("?")[0]
    return url.split("/")[-1].split("?")[0]

@app.route('/transcript', methods=['POST'])
def get_transcript():
    data = request.json
    video_url = data.get('url', '')
    
    if not video_url:
        return jsonify({"error": "URL não fornecida"}), 400

    try:
        video_id = extract_video_id(video_url).strip()
        print(f"Tentando processar ID: {video_id}")

        # Único método que seu sistema confirmou ter: api.list()
        api = YouTubeTranscriptApi()
        transcript_list = api.list(video_id)
        
        try:
            # Tenta pegar em português ou inglês
            transcript = transcript_list.find_transcript(['pt', 'en'])
        except:
            # Se não achar esses idiomas, pega a primeira disponível (qualquer uma)
            transcript = next(iter(transcript_list))

        raw_data = transcript.fetch()
        
        # Lógica para evitar o erro 'not subscriptable' (aceita objeto ou dicionário)
        full_text = ""
        for t in raw_data:
            if isinstance(t, dict):
                full_text += t.get('text', '') + " "
            else:
                full_text += getattr(t, 'text', '') + " "
        
        return jsonify({
            "status": "success",
            "video_id": video_id,
            "language": getattr(transcript, 'language', 'unknown'),
            "transcript": full_text.strip()
        })

    except Exception as e:
        print(f"Erro real no terminal: {str(e)}")
        return jsonify({"error": str(e)}), 500

if __name__ == '__main__':
    app.run(port=5000, debug=True)