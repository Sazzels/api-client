import { expect, test } from '@playwright/test';
import { cleanupEnvironment, cleanupRequest, setupEnvironment, setupRequest } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();
const requestSetupUUID = crypto.randomUUID();
const environmentSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID, RequestTypes.HTTP);
	await setupEnvironment(page, environmentSetupUUID);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
	await cleanupEnvironment(page, environmentSetupUUID);
});

test('request workflow', async ({ page }) => {
	const url = 'http://localhost:12346/echo';
	await page.locator('#request-method').selectOption('POST');
	await expect(page.locator('#request-method')).toHaveValue('POST');
	await page.locator('#request-url').fill(url);
	await expect(page.locator('#request-url')).toHaveValue(url);
	await expect(page.locator('#run-request')).not.toBeDisabled();
	await page.locator('#run-request').click();
	await expect(page.getByTestId('status-code')).toHaveText('200');

	// body
	await page.getByTestId('request-tabs').getByText('Body').click();
	await page.locator('#body-type').selectOption('json');
	await expect(page.getByTestId('json-body-error')).toBeVisible();
	await page.getByPlaceholder('> your body here <').fill('{"foo": "bar"}');
	await expect(page.getByTestId('json-body-error')).not.toBeVisible();

	// parameter
	await page.getByTestId('request-tabs').getByText('Parameter').click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url);
	await page.locator('#new-parameter-name').fill('hello');
	await page.locator('#new-parameter-value').fill('world');
	await page.getByLabel('save').click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url + '?hello=world');
	await page.getByTestId('request-parameters').getByLabel('delete').click();
	await expect(page.getByTestId('request-parameters')).toBeEmpty();

	// header
	await page.getByTestId('request-tabs').getByText('Header').click();
	await page.locator('#new-header-key').fill('foo');
	await page.locator('#new-header-value').fill('bar');
	await page.getByLabel('save').click();
	await expect(page.getByTestId('request-headers')).toHaveCount(1);
	await page.getByTestId('request-headers').getByLabel('delete').click();
	await expect(page.getByTestId('request-headers')).toBeEmpty();

	await page.locator('#run-request').click();
	await expect(page.getByTestId('payload')).toHaveText('{"foo": "bar"}');

	// select environment
	await page.locator('#current-environment').selectOption(environmentSetupUUID);
	await page.getByRole('button', { name: requestSetupUUID }).click();
	await page.getByTestId('request-tabs').getByText('Header').click();
	await expect(page.getByTestId('environment-request-header')).toHaveCount(2);

	await page.locator('#run-request').click();
	await expect(page.getByTestId('send-headers')).toContainText('Tom riddle');
	await expect(page.getByTestId('send-headers')).toContainText('Perry hotter');

	await page.getByTestId('request-tabs').getByText('Header').click();
	await page.getByTestId('environment-request-header').getByLabel('hide').first().click();
	await page.locator('#run-request').click();
	await expect(page.getByTestId('send-headers')).not.toContainText('Tom riddle');
	await expect(page.getByTestId('send-headers')).toContainText('Perry hotter');
});
