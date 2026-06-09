import { useRef, useState } from 'react';

interface UploadAreaProps {
  onAnalyze: (file: File) => void;
  isLoading: boolean;
}

export function UploadArea({ onAnalyze, isLoading }: UploadAreaProps) {
  const [file, setFile] = useState<File | null>(null);
  const inputRef = useRef<HTMLInputElement | null>(null);

  return (
    <section className="upload-card">
      <div className="upload-content">
        <div>
          <p className="upload-title">Upload subtitle file</p>
          <p className="upload-description">
            Choose a Japanese <strong>.srt</strong> file to analyze JLPT vocabulary density.
          </p>

          {file && (
            <p className="selected-file">
              Selected file: <strong>{file.name}</strong>
            </p>
          )}
        </div>

        <div className="upload-actions">
          <input
            ref={inputRef}
            className="hidden-file-input"
            type="file"
            accept=".srt"
            onChange={(e) => {
              if (e.target.files && e.target.files.length > 0) {
                setFile(e.target.files[0]);
              }
            }}
          />

          <button
            type="button"
            className="secondary-button"
            onClick={() => inputRef.current?.click()}
          >
            Choose file
          </button>

          <button
            type="button"
            className="primary-button"
            onClick={() => file && onAnalyze(file)}
            disabled={!file || isLoading}
          >
            {isLoading ? 'Analyzing...' : 'Analyze Subtitle'}
          </button>
        </div>
      </div>
    </section>
  );
}
