import { expect, test } from '@playwright/test';
import { cleanupCollections, setupCollections } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();

test.beforeEach('collection and project setup', async ({ page }) => {
	await setupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test.afterEach('collection and project cleanup', async ({ page }) => {
	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test('http request workflow', async ({ page }) => {
	const httpRequestUUID = crypto.randomUUID();

	await page.locator('#new-request').fill(httpRequestUUID);
	await page.locator('#request-type').selectOption(RequestTypes.HTTP);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('button', { name: httpRequestUUID })).toBeVisible();

	//update
	await page.getByTestId('request').filter({ hasText: httpRequestUUID }).getByLabel('edit').click();

	const otherHttpRequestUUID = crypto.randomUUID();
	await page.getByTestId('requests').getByRole('textbox').fill(otherHttpRequestUUID);
	await page.getByTestId('requests').getByLabel('save').click();

	// delete
	await page.getByTestId('request').filter({ hasText: otherHttpRequestUUID }).getByLabel('delete').click();

	await expect(page.getByTestId('request').getByRole('button', { name: otherHttpRequestUUID })).not.toBeVisible();
});

test('websocket request workflow', async ({ page }) => {
	const websocketRequestUUID = crypto.randomUUID();

	await page.locator('#new-request').fill(websocketRequestUUID);
	await page.locator('#request-type').selectOption(RequestTypes.WEBSOCKET);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('button', { name: websocketRequestUUID })).toBeVisible();

	//update
	await page.getByTestId('request').filter({ hasText: websocketRequestUUID }).getByLabel('edit').click();

	const otherWebsocketRequestUUID = crypto.randomUUID();
	await page.getByTestId('requests').getByRole('textbox').fill(otherWebsocketRequestUUID);
	await page.getByTestId('requests').getByLabel('save').click();

	// delete
	await page.getByTestId('request').filter({ hasText: otherWebsocketRequestUUID }).getByLabel('delete').click();

	await expect(page.getByTestId('request').getByRole('button', { name: otherWebsocketRequestUUID })).not.toBeVisible();
});
