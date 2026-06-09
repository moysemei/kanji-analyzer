import { useState } from 'react';
import type { AnalysisResult, VocabularyItem } from '../types';

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
    <section className="result-panel">
      <div className="result-header">
        <h2 style={{ margin: 0 }}>Episode Summary</h2>

        <button className="download-button" onClick={onDownload}>
          Download CSV for Anki
        </button>
      </div>

      <div className="total-card">
        <span className="total-label">Total Unique Words</span>
        <h3 className="total-value">{result.stats.totalWords}</h3>
      </div>

      <h3 className="section-title">Density by Level (JLPT)</h3>

      <label className="toggle-label">
        <input
          type="checkbox"
          checked={showUnknown}
          onChange={(event) => setShowUnknown(event.target.checked)}
        />
        Show unknown words
      </label>

      <div className="level-grid">
        {jlptLevels.map((level) => (
          <div key={level} className="level-card">
            <div className="level-card-header">
              <strong className="level-name">{level}</strong>

              <span className="level-count">
                {result.stats.levelCount[level]} words
              </span>
            </div>

            <div className="progress-track">
              <div
                className="progress-fill"
                style={{ width: `${result.stats.density[level]}%` }}
              />
            </div>

            <p className="level-density">
              {result.stats.density[level].toFixed(1)}% of text
            </p>
          </div>
        ))}
      </div>

      <h3 className="section-title">Extracted Vocabulary</h3>

      <div className="vocabulary-panel">
        {visibleLevels.map((level) => {
          const items = groupedVocabulary[level] ?? [];

          if (items.length === 0) {
            return null;
          }

          return (
            <section key={level}>
              <h4 className="vocabulary-section-title">
                {level} ({items.length})
              </h4>

              <div className="vocabulary-list">
                {items.map((item, index) => (
                  <span className="vocabulary-chip" key={`${item.word}-${index}`}>
                    {item.word}

                    {item.reading && item.reading !== item.word && (
                      <small className="vocabulary-reading">
                        （{item.reading}）
                      </small>
                    )}
                  </span>
                ))}
              </div>
            </section>
          );
        })}
      </div>
    </section>
  );
}
