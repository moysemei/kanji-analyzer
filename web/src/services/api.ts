import type { AnalysisResult } from '../types';

export const analyzeSubtitle = async (file: File): Promise<AnalysisResult> => {
	const formData = new FormData();
	formData.append('subtitle', file);

	const response = await fetch('http://localhost:8080/api/analyze', {
		method: 'POST',
		body: formData,
	});

	if (!response.ok) {
		throw new Error(`Server error: ${response.statusText}`);
	}

	return response.json();
}
