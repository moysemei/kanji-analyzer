export interface VocabularyItem {
  word: string;
  reading?: string;
  level: string;
}

export interface AnalysisResult {
  stats: {
    totalWords: number;
    levelCount: Record<string, number>;
    density: Record<string, number>;
  };
  vocabulary: VocabularyItem[];
}
