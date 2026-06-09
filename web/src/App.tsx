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

    const csvContent = [
  'word,level',
  ...result.vocabulary.map((item) => `${item.word},${item.level}`)
].join('\n');
    const blob = new Blob([new Uint8Array([0xEF, 0xBB, 0xBF]), csvContent], { 
      type: 'text/csv;charset=utf-8;' 
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
    <div style={{ padding: '2rem', fontFamily: 'sans-serif' }}>
      <h1>Kanji Analyzer</h1>
      <p>Upload a subtitle file (.srt) to extract JLPT vocabulary.</p>
      
      {/* Component 1: Upload Area */}
      <UploadArea onAnalyze={handleAnalyze} isLoading={isLoading} />

      {/* Component 2: Result Panel (only shows if there is results) */}
      {result && <ResultPanel result={result} onDownload={downloadCSV} />}
    </div>
  );
}

export default App;
