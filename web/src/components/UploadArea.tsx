import { useState } from 'react';

interface UploadAreaProps {
  onAnalyze: (file: File) => void;
  isLoading: boolean;
}

export function UploadArea({ onAnalyze, isLoading }: UploadAreaProps) {
  const [file, setFile] = useState<File | null>(null);

  return (
    <div style={{ marginTop: '2rem' }}>
      <input 
        type="file" 
        accept=".srt"
        onChange={(e) => {
          if (e.target.files && e.target.files.length > 0) {
            setFile(e.target.files[0]);
          }
        }}
      />
      
      {file && (
        <div style={{ marginTop: '1rem' }}>
          <p>Selected file: <strong>{file.name}</strong></p>
          <button 
            onClick={() => onAnalyze(file)} 
            disabled={isLoading}
            style={{ padding: '0.5rem 1rem', cursor: 'pointer' }}
          >
            {isLoading ? 'Analyzing...' : 'Analyze Subtitle'}
          </button>
        </div>
      )}
    </div>
  );
}
