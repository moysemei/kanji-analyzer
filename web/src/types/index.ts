export interface AnalysisResult {
	stats: {
		TotalWords: number;
		LevelCount: Record<string, number>;
		Density: Record<string, number>;
	};
	vocabulary: string[];
}
