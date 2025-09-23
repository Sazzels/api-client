import { expect, test } from '@playwright/test';

test('environment handling', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await page.getByTestId('open-environment').click();
	const environmentUUID = crypto.randomUUID();
	await page.locator('#new-environment').fill(environmentUUID);
	await expect(page.locator('#new-environment')).toHaveValue(environmentUUID);
	await page.locator('#create-new-environment').click();
	await expect(page.locator('#new-environment')).toBeEmpty();
	await expect(page.getByTestId('environments').getByRole('button', { name: environmentUUID })).toBeVisible();

	// update
	await page.getByTestId('environment').filter({ hasText: environmentUUID }).getByLabel('edit').click();
	const otherEnvironmentUUID = crypto.randomUUID();
	await page.getByTestId('environments').getByRole('textbox').fill(otherEnvironmentUUID);
	await page.getByTestId('environments').getByLabel('save').click();

	// add header
	await page.getByTestId('environments').getByRole('button', { name: otherEnvironmentUUID }).click();
	await page.locator('#new-header-key').fill('tom');
	await page.locator('#new-header-value').fill('riddle');
	await page.getByLabel('save').click();
	await expect(page.locator('#new-header-key')).toBeEmpty();
	await expect(page.locator('#new-header-value')).toBeEmpty();
	await expect(page.getByTestId('environment-header')).toHaveCount(1);
	await page.locator('#new-header-key').fill('perry');
	await page.locator('#new-header-value').fill('hotter');
	await page.getByLabel('save').click();
	await expect(page.getByTestId('environment-header')).toHaveCount(2);

	// delete header
	await page.getByTestId('environment-header').getByLabel('delete').last().click();
	await expect(page.getByTestId('environment-header')).toHaveCount(1);
	// update header
	await page.getByTestId('environment-headers').getByRole('textbox').first().fill('lord');
	await page.getByTestId('environment-headers').getByRole('textbox').last().fill('voldemort');

	// delete
	await page.getByTestId('environment').filter({ hasText: otherEnvironmentUUID }).getByLabel('delete').click();
	await expect(page.getByTestId('environment').filter({ hasText: otherEnvironmentUUID })).toHaveCount(0);
});
