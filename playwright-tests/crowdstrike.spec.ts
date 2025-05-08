import { test, expect } from '@playwright/test';

test('CrowdStrike interactions', async ({ page }) => {
  // Navigate to CrowdStrike homepage
  await page.goto('/en-us/');

  // Handle any popup dialogs
  try {
    await page.getByRole('button', { name: 'Close' }).click();
  } catch (e) {
    console.log('First close button not found');
  }

  try {
    await page.getByRole('button', { name: 'Close', exact: true }).click();
  } catch (e) {
    console.log('Second close button not found');
  }

  // Click on the search icon
  await page.getByRole('button', { name: 'Search Icon' }).click();

  // Type search query
  await page.getByRole('searchbox', { name: 'Search field' }).fill('blash');

  // Click on the search results text
  await page.getByText('Results for blash').click();

  // Wait a bit to ensure all requests are captured
  await page.waitForTimeout(2000);
}); 