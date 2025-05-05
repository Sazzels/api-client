import { expect, test } from '@playwright/test';
import { cleanupRequest, setupRequest } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();
const requestSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID, RequestTypes.HTTP);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
});

test('request workflow with clipboard functionality', async ({ page }) => {
	const url = 'https://jsonplaceholder.typicode.com/posts';
	await page.locator('#request-method').selectOption('GET');
	await expect(page.locator('#request-method')).toHaveValue('GET');
	await page.locator('#request-url').fill(url);
	await page.locator('#run-request-icon').click();

	await page.waitForSelector('.text-text-highlight:has-text("Status-Code:")');
	// const statusCodeSection = page.locator('div.flex.justify-between', { hasText: 'Status-Code:' });
	// await expect(statusCodeSection).toBeVisible();

	const statusCodeClipboardButton = page.locator('#status-code-clipboard-btn');
	await statusCodeClipboardButton.waitFor({ state: 'visible' });
	await statusCodeClipboardButton.click();
	await page.waitForTimeout(500); // FIXME: timeout for clipboard action
	const clipboardContent = await page.evaluate(() => navigator.clipboard.readText());

	// Log for debugging
	console.log('Clipboard content:', clipboardContent);
});
