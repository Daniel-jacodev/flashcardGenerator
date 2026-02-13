import { useState } from "react";
import axios from "axios";
import { FileText, Youtube, Download, Loader2, Play } from "lucide-react";

function App() {
  const [file, setFile] = useState(null);
  const [url, setUrl] = useState("");
  const [flashcards, setFlashcards] = useState("");
  const [loading, setLoading] = useState(false);
  const [mode, setMode] = useState("file"); // 'file' ou 'url'

  const handleGenerate = async () => {
    setLoading(true);
    const formData = new FormData();

    if (mode === "file") {
      if (!file) return alert("Selecione um arquivo!");
      formData.append("file", file);
    } else {
      if (!url) return alert("Cole uma URL do YouTube!");
      formData.append("url", url);
    }

    try {
      const response = await axios.post(
        "http://localhost:8080/generate",
        formData
      );
      setFlashcards(response.data.flashcards);
    } catch (error) {
      alert("Erro ao gerar flashcards. Verifique o backend.");
    } finally {
      setLoading(false);
    }
  };

  const downloadCSV = () => {
    const blob = new Blob([flashcards], { type: "text/csv" });
    const urlBlob = window.URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = urlBlob;
    a.download = `flashcards_${Date.now()}.csv`;
    a.click();
  };

  return (
    <div className="min-h-screen bg-slate-950 text-slate-100 p-4 md:p-8 font-sans">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <header className="text-center mb-12">
          <h1 className="text-4xl font-extrabold bg-gradient-to-r from-blue-400 to-emerald-400 bg-clip-text text-transparent">
            Flashcard Generator AI
          </h1>
          <p className="text-slate-400 mt-2">Estude com eficiência usando IA</p>
        </header>

        {/* Seletor de Modo */}
        <div className="flex justify-center gap-4 mb-8">
          <button
            onClick={() => setMode("file")}
            className={`flex items-center gap-2 px-6 py-2 rounded-full transition ${
              mode === "file"
                ? "bg-blue-600"
                : "bg-slate-800 hover:bg-slate-700"
            }`}
          >
            <FileText size={20} /> Arquivo
          </button>
          <button
            onClick={() => setMode("url")}
            className={`flex items-center gap-2 px-6 py-2 rounded-full transition ${
              mode === "url" ? "bg-red-600" : "bg-slate-800 hover:bg-slate-700"
            }`}
          >
            <Youtube size={20} /> YouTube
          </button>
        </div>

        {/* Card Principal */}
        <div className="bg-slate-900 border border-slate-800 p-8 rounded-2xl shadow-xl">
          {mode === "file" ? (
            <div className="text-center">
              <input
                type="file"
                id="pdf"
                className="hidden"
                onChange={(e) => setFile(e.target.files[0])}
                accept=".pdf,.mp3,.m4a"
              />
              <label
                htmlFor="pdf"
                className="cursor-pointer block border-2 border-dashed border-slate-700 rounded-xl p-12 hover:border-blue-500 transition"
              >
                <FileText className="mx-auto mb-4 text-slate-500" size={48} />
                <p className="text-slate-300">
                  {file ? file.name : "Arraste ou selecione PDF/MP3"}
                </p>
              </label>
            </div>
          ) : (
            <div className="space-y-4">
              <label className="block text-sm font-medium text-slate-400">
                Link do vídeo no YouTube
              </label>
              <input
                type="text"
                placeholder="https://www.youtube.com/watch?v=..."
                className="w-full bg-slate-950 border border-slate-800 rounded-lg px-4 py-3 focus:ring-2 focus:ring-red-500 outline-none"
                value={url}
                onChange={(e) => setUrl(e.target.value)}
              />
            </div>
          )}

          <button
            onClick={handleGenerate}
            disabled={loading}
            className="w-full mt-8 bg-emerald-600 hover:bg-emerald-500 disabled:bg-slate-700 py-4 rounded-xl font-bold text-lg flex justify-center items-center gap-2 transition"
          >
            {loading ? (
              <>
                <Loader2 className="animate-spin" /> Gerando Cards...
              </>
            ) : (
              <>
                <Play /> Começar Mágica
              </>
            )}
          </button>
        </div>

        {/* Resultados */}
        {flashcards && (
          <div className="mt-12 animate-in fade-in duration-500">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-bold">Cards Prontos</h2>
              <button
                onClick={downloadCSV}
                className="flex items-center gap-2 text-emerald-400 hover:text-emerald-300 font-medium"
              >
                <Download size={20} /> Baixar .CSV para Anki
              </button>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {flashcards
                .split("\n")
                .filter((l) => l.includes(";"))
                .map((card, i) => (
                  <div
                    key={i}
                    className="bg-slate-900 p-6 rounded-xl border-l-4 border-blue-500"
                  >
                    <p className="text-xs text-blue-400 uppercase font-bold tracking-widest mb-1">
                      Frente
                    </p>
                    <p className="text-slate-200 mb-4">{card.split(";")[0]}</p>
                    <p className="text-xs text-emerald-400 uppercase font-bold tracking-widest mb-1">
                      Verso
                    </p>
                    <p className="text-slate-400">{card.split(";")[1]}</p>
                  </div>
                ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
