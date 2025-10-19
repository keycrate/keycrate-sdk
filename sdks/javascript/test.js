import { configurate } from './dist/index.js';

async function main() {
  const host = process.argv[2] || 'http://127.0.0.1:8787';
  const appId = process.argv[3] || '57d87dfa-18a6-4eed-9074-f37418067c47';

  console.log(`Connecting to: ${host}`);
  console.log(`App ID: ${appId}\n`);

  const client = configurate(host, appId);

  // Test 1: Authenticate with license
  console.log('Test 1: Authenticate with license key');
  const authResult1 = await client.authenticate({
    license: 'test-license-key'
  });
  console.log('Result:', authResult1);
  console.log('');

  // Test 2: Authenticate with username/password
  console.log('Test 2: Authenticate with username/password');
  const authResult2 = await client.authenticate({
    username: 'testuser',
    password: 'testpass'
  });
  console.log('Result:', authResult2);
  console.log('');

  // Test 3: Missing authentication method
  console.log('Test 3: Missing authentication method (should fail)');
  const authResult3 = await client.authenticate({});
  console.log('Result:', authResult3);
  console.log('');

  // Test 4: Register
  console.log('Test 4: Register credentials');
  const registerResult = await client.register({
    license: 'test-license',
    username: 'newuser',
    password: 'newpass'
  });
  console.log('Result:', registerResult);
}

main().catch(console.error);