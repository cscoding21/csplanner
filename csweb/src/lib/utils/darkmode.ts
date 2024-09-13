/**
 * A function to determine if the user is browsing in dark mode, as persisted in local storage
 * @returns true if in dark mode.  False otherwise
 */
export const isDarkMode = (): boolean => {
	const dm = localStorage.getItem('color-theme');

	return dm === 'dark';
};
