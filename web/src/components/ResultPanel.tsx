import type { AnalysisResult, VocabularyItem } from '../types';
import { useState } from 'react';

interface ResultPanelProps {
  result: AnalysisResult;
  onDownload: () => void;
}

const levelOrder = ['N5', 'N4', 'N3', 'N2', 'N1', 'Unknown'];

export function ResultPanel({ result, onDownload }: ResultPanelProps) {
  const [showUnknown, setShowUnknown] = useState(false);

  const allJlptLevels = Object.keys(result.stats.density).sort().reverse();

  const jlptLevels = showUnknown
    ? allJlptLevels
    : allJlptLevels.filter((level) => level !== 'Unknown');

  const groupedVocabulary = result.vocabulary.reduce<Record<string, VocabularyItem[]>>(
    (groups, item) => {
      if (!groups[item.level]) {
        groups[item.level] = [];
      }

      groups[item.level].push(item);

      return groups;
    },
    {}
  );

  const visibleLevels = showUnknown
    ? levelOrder
    : levelOrder.filter((level) => level !== 'Unknown');
  return (
    <div style={{ marginTop: '2rem', textAlign: 'left' }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', borderBottom: '2px solid #eee', paddingBottom: '1rem', marginBottom: '1.5rem' }}>
        <h2 style={{ margin: 0, color: '#333' }}>Episode Summary</h2>

        <button
          onClick={onDownload}
          style={{ padding: '0.5rem 1rem', backgroundColor: '#28a745', color: 'white', border: 'none', borderRadius: '4px', cursor: 'pointer', fontWeight: 'bold' }}
        >
          Download CSV for Anki
        </button>
      </div>

      <div style={{ display: 'flex', gap: '1rem', marginBottom: '2rem' }}>
        <div style={{ background: '#f8f9fa', padding: '1rem', borderRadius: '8px', flex: 1, border: '1px solid #e9ecef' }}>
          <span style={{ fontSize: '0.9rem', color: '#6c757d' }}>Total Unique Words</span>

          <h3 style={{ margin: '0.5rem 0 0 0', fontSize: '2rem', color: '#2b2b2b' }}>
            {result.stats.totalWords}
          </h3>
        </div>
      </div>

      <h3 style={{ color: '#495057', marginBottom: '1rem' }}>Density by Level (JLPT)</h3>

	  <label style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '1rem', color: '#495057' }}>
  <input
    type="checkbox"
    checked={showUnknown}
    onChange={(event) => setShowUnknown(event.target.checked)}
  />
  Show unknown words
</label>      <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(150px, 1fr))', gap: '1rem', marginBottom: '2rem' }}>
        {jlptLevels.map((level) => (
          <div key={level} style={{ background: '#ffffff', padding: '1rem', borderRadius: '8px', border: '1px solid #dee2e6', boxShadow: '0 2px 4px rgba(0,0,0,0.05)' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '0.5rem' }}>
              <strong style={{ fontSize: '1.2rem', color: '#007bff' }}>{level}</strong>

              <span style={{ fontSize: '0.85rem', color: '#6c757d', fontWeight: 'bold' }}>
                {result.stats.levelCount[level]} words
              </span>
            </div>

            <div style={{ background: '#e9ecef', height: '8px', borderRadius: '4px', overflow: 'hidden' }}>
              <div style={{ background: '#007bff', height: '100%', width: `${result.stats.density[level]}%` }} />
            </div>

            <p style={{ margin: '0.5rem 0 0 0', fontSize: '0.85rem', textAlign: 'right', color: '#495057' }}>
              {result.stats.density[level].toFixed(1)}% of text
            </p>
          </div>
        ))}
      </div>

      <h3 style={{ color: '#495057', marginBottom: '1rem' }}>Extracted Vocabulary</h3>

<label style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '1rem', color: '#495057' }}>
  <input
    type="checkbox"
    checked={showUnknown}
    onChange={(event) => setShowUnknown(event.target.checked)}
  />
  Show unknown words
</label>

<div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem', background: '#f8f9fa', padding: '1.5rem', borderRadius: '8px', border: '1px solid #e9ecef', maxHeight: '420px', overflowY: 'auto' }}>
       {visibleLevels.map((level) => {
          const items = groupedVocabulary[level] ?? [];

          if (items.length === 0) {
            return null;
          }

          return (
            <section key={level}>
              <h4 style={{ margin: '0 0 0.75rem 0', color: '#343a40' }}>
                {level} ({items.length})
              </h4>

              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '0.5rem' }}>
                {items.map((item, index) => (
                  <span
                    key={`${item.word}-${index}`}
                    style={{ background: '#ffffff', padding: '0.25rem 0.75rem', borderRadius: '16px', border: '1px solid #ced4da', fontSize: '0.95rem', color: '#343a40' }}
                  >
                    {item.word}
                  </span>
                ))}
              </div>
            </section>
          );
        })}
		      </div>
    </div>
  );
}
