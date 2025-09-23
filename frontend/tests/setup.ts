import { expect, type Page } from '@playwright/test';
import type { RequestTypes } from '$lib/enums/RequestTypes.ts';

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

	await page.locator('#new-collection').fill(collectionSetupUUID);
	await page.locator('#create-new-collection').click();

	await expect(page.locator('#new-collection')).toBeEmpty();
	await expect(page.getByTestId('collections').getByRole('button', { name: collectionSetupUUID })).toBeVisible();
	await page.getByTestId('collections').getByRole('button', { name: collectionSetupUUID }).click();

	await expect(page).toHaveTitle('Api-Client :: Collection Overview');
	await expect(page.getByRole('button', { name: collectionSetupUUID })).toBeVisible();
}

export async function cleanupCollections(page: Page, projectSetupUUID: string, collectionsSetupUUID: string) {
	await navigateToProject(page);
	await page.getByTestId('projects').getByRole('button', { name: projectSetupUUID }).click();
	await expect(page).toHaveTitle('Api-Client :: Project Details');

	await expect(page.getByRole('button', { name: projectSetupUUID })).toBeVisible();
	await page.getByTestId('collection').filter({ hasText: collectionsSetupUUID }).getByLabel('delete').click();

	await cleanupProjects(page, projectSetupUUID);
}

export async function setupRequest(
	page: Page,
	projectSetupUUID: string,
	collectionSetupUUID: string,
	requestSetupUUID: string,
	requestType: RequestTypes,
) {
	await setupCollections(page, projectSetupUUID, collectionSetupUUID);

	await page.locator('#request-type').selectOption(requestType);
	await page.locator('#new-request').fill(requestSetupUUID);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
	await expect(page.getByRole('button', { name: requestSetupUUID })).toBeVisible();

	await page.getByTestId('requests').getByRole('button', { name: requestSetupUUID }).click();
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

	await page.getByTestId('requests').filter({ hasText: requestSetupUUID }).getByLabel('delete').click();

	await expect(page.getByRole('button', { name: requestSetupUUID })).not.toBeVisible();

	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
}

export async function setupEnvironment(page: Page, environmentSetupUUID: string) {
	await page.getByTestId('open-environment').click();
	await page.locator('#new-environment').fill(environmentSetupUUID);
	await expect(page.locator('#new-environment')).toHaveValue(environmentSetupUUID);
	await page.locator('#create-new-environment').click();
	await expect(page.locator('#new-environment')).toBeEmpty();
	await expect(page.getByTestId('environments').getByRole('button', { name: environmentSetupUUID })).toBeVisible();

	await page.getByTestId('environments').getByRole('button', { name: environmentSetupUUID }).click();
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

	await page.getByRole('dialog').press('Escape');
}

export async function cleanupEnvironment(page: Page, environmentSetupUUID: string) {
	await page.getByTestId('open-environment').click();
	await page.getByTestId('environment').filter({ hasText: environmentSetupUUID }).getByLabel('delete').click();
	await expect(page.getByTestId('environment').filter({ hasText: environmentSetupUUID })).toHaveCount(0);
	await page.getByRole('dialog').press('Escape');
}
