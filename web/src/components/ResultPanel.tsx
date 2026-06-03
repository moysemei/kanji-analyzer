import { AnalysisResult } from '../types';

interface ResultPanelProps {
  result: AnalysisResult;
  onDownload: () => void;
}

export function ResultPanel({ result, onDownload }: ResultPanelProps) {
  return (
    <div style={{ marginTop: '2rem', textAlign: 'left', background: '#f5f5f5', padding: '1rem', borderRadius: '8px' }}>
      <h3 style={{ color: '#333', display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
        <span>Analysis Complete</span>
        <button 
          onClick={onDownload}
          style={{ 
            padding: '0.5rem 1rem', 
            backgroundColor: '#28a745', 
            color: 'white', 
            border: 'none', 
            borderRadius: '4px', 
            cursor: 'pointer',
            fontWeight: 'bold'
          }}
        >
        Download CSV for Anki
        </button>
      </h3>
      
      <pre style={{ color: '#333', overflowX: 'auto' }}>
        {JSON.stringify(result, null, 2)}
      </pre>
    </div>
  );
}
