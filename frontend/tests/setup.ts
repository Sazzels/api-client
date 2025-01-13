import { expect, type Page } from '@playwright/test';

export async function navigateToProject(page: Page): Promise<void> {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.getByRole('button', { name: 'to project overview' }).click();
	await expect(page).toHaveTitle('Api-Client :: Project Overview');
}

export async function setupProjects(page: Page, projectSetupUUID: string): Promise<void> {
	await navigateToProject(page);

	await page.locator('#new-project').fill(projectSetupUUID);
	await page.locator('#create-new-project').click();
	await page.getByTestId('projects').getByRole('button', { name: projectSetupUUID }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Details');
	await expect(page.getByRole('button', { name: projectSetupUUID })).toBeVisible();
}

export async function cleanupProjects(page: Page, projectSetupUUID: string) {
	await navigateToProject(page);

	await page.getByTestId('project').filter({ hasText: projectSetupUUID }).getByLabel('delete').click();
	await expect(page.getByTestId('projects').getByRole('button', { name: projectSetupUUID })).not.toBeVisible();
}

export async function setupCollections(page: Page, projectSetupUUID: string, collectionSetupUUID: string) {
	await setupProjects(page, projectSetupUUID);
	await page.getByRole('textbox', { name: 'insert new collection name' }).fill(collectionSetupUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new collection name' })).toBeEmpty();
	await expect(page.getByTestId('collections').getByRole('button', { name: collectionSetupUUID })).toBeVisible();
	await page.getByTestId('collections').getByRole('button', { name: collectionSetupUUID }).click();
	await expect(page.getByRole('button', { name: collectionSetupUUID })).toBeVisible();
}

export async function cleanupCollections(page: Page, projectSetupUUID: string, collectionsSetupUUID: string) {
	await navigateToProject(page);
	await page.getByTestId('projects').getByRole('button', { name: projectSetupUUID }).click();
	await expect(page.getByRole('button', { name: projectSetupUUID })).toBeVisible();
	await page
		.getByRole('listitem')
		.filter({ hasText: collectionsSetupUUID })
		.getByRole('button', { name: 'delete' })
		.click();

	await expect(page.getByRole('listitem').getByRole('button', { name: collectionsSetupUUID })).not.toBeVisible();
	await cleanupProjects(page, projectSetupUUID);
}

export async function setupRequest(
	page: Page,
	projectSetupUUID: string,
	collectionSetupUUID: string,
	requestSetupUUID: string,
) {
	await setupCollections(page, projectSetupUUID, collectionSetupUUID);

	await page.getByRole('textbox', { name: 'insert new request name' }).fill(requestSetupUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new request name' })).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('listitem').filter({ hasText: requestSetupUUID })).toBeVisible();
	await page.getByTestId('requests').getByRole('listitem').filter({ hasText: requestSetupUUID }).click();
}

export async function cleanupRequest(
	page: Page,
	projectSetupUUID: string,
	collectionsSetupUUID: string,
	requestSetupUUID: string,
) {
	await navigateToProject(page);
	await page.getByTestId('projects').getByRole('button', { name: projectSetupUUID }).click();
	await expect(page.getByRole('button', { name: projectSetupUUID })).toBeVisible();
	await page.getByTestId('collections').getByRole('button', { name: collectionsSetupUUID }).click();
	await expect(page.getByRole('button', { name: collectionsSetupUUID })).toBeVisible();
	await page
		.getByRole('listitem')
		.filter({ hasText: requestSetupUUID })
		.getByRole('button', { name: 'delete' })
		.click();

	await expect(
		page.getByTestId('requests').getByRole('listitem').filter({ hasText: requestSetupUUID }),
	).not.toBeVisible();

	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
}
