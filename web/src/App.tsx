import { useState } from 'react';
import './App.css';
import type { AnalysisResult } from './types';
import { analyzeSubtitle } from './services/api';
import { UploadArea } from './components/UploadArea';
import { ResultPanel } from './components/ResultPanel';

function App() {
  const [isLoading, setIsLoading] = useState(false);
  const [result, setResult] = useState<AnalysisResult | null>(null);

  const handleAnalyze = async (file: File) => {
    setIsLoading(true);
    setResult(null); 

    try {
      const data = await analyzeSubtitle(file);
      setResult(data);
    } catch (error) {
      console.error("Communication failed:", error);
      alert("Failed to communicate with the Go backend. Is the server running?");
    } finally {
      setIsLoading(false);
    }
  };

  const downloadCSV = () => {
  if (!result) return;

  const escapeCSV = (value: string) => {
    const escaped = value.replace(/"/g, '""');
    return `"${escaped}"`;
  };

  const csvContent = [
    'word,reading,level',
    ...result.vocabulary.map((item) =>
      `${escapeCSV(item.word)},${escapeCSV(item.reading ?? '')},${escapeCSV(item.level)}`
    ),
  ].join('\n');

  const blob = new Blob([new Uint8Array([0xef, 0xbb, 0xbf]), csvContent], {
    type: 'text/csv;charset=utf-8;',
  });

  const url = URL.createObjectURL(blob);
  const link = document.createElement('a');
  link.href = url;
  link.setAttribute('download', 'anki_deck.csv');
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
};
  return (
  <main className="app-shell">
    <section className="hero">
      <div>
        <div className="hero-badge">字幕 → JLPT vocabulary</div>
        <h1 className="hero-title">Kanji Analyzer</h1>
        <p className="hero-subtitle">
          Upload a Japanese subtitle file and discover which JLPT vocabulary levels appear in the episode.
        </p>
      </div>

      <div className="hero-card">
        <p className="hero-kanji">漢字</p>
        <p className="hero-card-text">Analyze, classify and export vocabulary for Anki.</p>
      </div>
    </section>

    <UploadArea onAnalyze={handleAnalyze} isLoading={isLoading} />

    {result && <ResultPanel result={result} onDownload={downloadCSV} />}
  </main>
);
}

export default App;
