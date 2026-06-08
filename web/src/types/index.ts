export interface AnalysisResult {
	stats: {
		totalWords: number;
		levelCount: Record<string, number>;
		density: Record<string, number>;
	};
	vocabulary: string[];
}
