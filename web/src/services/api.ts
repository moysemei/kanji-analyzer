import type { AnalysisResult } from '../types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080';

export const analyzeSubtitle = async (file: File): Promise<AnalysisResult> => {
	const formData = new FormData();
	formData.append('subtitle', file);

	const response = await fetch(`${API_BASE_URL}/api/analyze`, {
		method: 'POST',
		body: formData,
	});

	if (!response.ok) {
		throw new Error(`Server error: ${response.statusText}`);
	}

	return response.json();
}
