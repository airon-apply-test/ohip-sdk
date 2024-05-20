import { RequestContext } from './oauth';

type OhipCashierMapping = {
  [ohipUsername: string]: string;
};

interface OhipJWT {
  sub: string;
}

export class OhipCredentialsProvider {
  usernameToCashierIdMapping: OhipCashierMapping;
  keyToReplaceWithCashierId: string;

  constructor({
    cashierMapping,
    keyToReplaceWithCashierId,
  }: {
    cashierMapping: OhipCashierMapping;
    keyToReplaceWithCashierId: string;
  }) {
    this.keyToReplaceWithCashierId = keyToReplaceWithCashierId;
    this.usernameToCashierIdMapping = cashierMapping;
  }

  async cashierMappingMiddleware(
    context: RequestContext,
  ): Promise<RequestContext> {
    const headersString = JSON.stringify(context.init.headers);
    const bodyString = JSON.stringify(context.init.body);
    const mustReplaceHeaders = headersString.includes(
      this.keyToReplaceWithCashierId,
    );
    const mustReplaceBody = bodyString.includes(this.keyToReplaceWithCashierId);
    if (!mustReplaceHeaders && !mustReplaceBody) return context;
    // @ts-ignore
    const token = context.init.headers?.Authorization?.split(' ')[1];
    const sub = getSubFromOhipAccessToken(token);
    if (!sub) throw new Error('invalid access token or sub missing');
    const cashierId = this.usernameToCashierIdMapping[sub];
    if (!cashierId)
      throw new Error('cashier id not found for sub specified in token');
    return {
      ...context,
      init: {
        ...context.init,
        body: mustReplaceBody
          ? JSON.parse(
              bodyString.replace(this.keyToReplaceWithCashierId, cashierId),
            )
          : context.init.body,
        headers: mustReplaceHeaders
          ? JSON.parse(
              headersString.replace(this.keyToReplaceWithCashierId, cashierId),
            )
          : context.init.headers,
      },
    };
  }
}

function getSubFromOhipAccessToken(token: string): string | null {
  const parts: string[] = token.split('.');
  if (parts.length < 2) {
    return null;
  }
  const payload: string = base64UrlDecode(parts[1]);
  let claims: OhipJWT;
  try {
    claims = JSON.parse(payload);
  } catch (error) {
    return null;
  }
  return claims.sub;
}

function base64UrlDecode(input: string): string {
  const base64 = input.replace(/-/g, '+').replace(/_/g, '/');
  const buffer = Buffer.from(base64, 'base64');
  return buffer.toString('utf-8');
}
