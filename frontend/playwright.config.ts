import { defineConfig, devices } from '@playwright/test';

export default defineConfig({
	webServer: {
		command: 'cd .. && make start-test-environment',
		timeout: 240000,
		port: 34115,
		reuseExistingServer: true,
		stderr: 'ignore',
		gracefulShutdown: { signal: 'SIGTERM', timeout: 5000 },
	},
	testDir: 'tests',
	testMatch: /(.+\.)?(test|spec)\.[jt]s/,
	projects: [
		/* Test against desktop browsers */
		{
			name: 'chromium',
			use: {
				...devices['Desktop Chrome'],
				permissions: ['clipboard-read', 'clipboard-write'],
			},
		},
		{
			name: 'firefox',
			use: {
				...devices['Desktop Firefox'],
				launchOptions: {
					firefoxUserPrefs: {
						'dom.events.testing.asyncClipboard': true,
						'dom.events.asyncClipboard.readText': true,
						'dom.events.asyncClipboard.clipboardItem': true,
						'dom.events.asyncClipboard.writeText': true,
						'permissions.default.clipboard-read': 1,
						'permissions.default.clipboard-write': 1,
					},
				},
			},
		},
		// {
		// 	name: 'webkit',
		// 	use: { ...devices['Desktop Safari'] },
		// },
		// /* Test against mobile viewports. */
		// {
		// 	name: 'Mobile Chrome',
		// 	use: { ...devices['Pixel 5'] }
		// },
		// {
		// 	name: 'Mobile Safari',
		// 	use: { ...devices['iPhone 12'] }
		// },
		// /* Test against branded browsers. */
		// {
		// 	name: 'Google Chrome',
		// 	use: { ...devices['Desktop Chrome'], channel: 'chrome' } // or 'chrome-beta'
		// },
		// {
		// 	name: 'Microsoft Edge',
		// 	use: { ...devices['Desktop Edge'], channel: 'msedge' } // or 'msedge-dev'
		// }
	],
});
