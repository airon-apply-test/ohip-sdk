/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { KeywordType } from './KeywordType';
import {
    KeywordTypeFromJSON,
    KeywordTypeFromJSONTyped,
    KeywordTypeToJSON,
} from './KeywordType';

/**
 * List of customer keywords.
 * @export
 * @interface ProfileTypeKeywords
 */
export interface ProfileTypeKeywords {
    /**
     * Collection of keywords attached to the profile.
     * @type {Array<KeywordType>}
     * @memberof ProfileTypeKeywords
     */
    keyword?: Array<KeywordType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ProfileTypeKeywords
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ProfileTypeKeywords
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ProfileTypeKeywords
     */
    count?: number;
}

/**
 * Check if a given object implements the ProfileTypeKeywords interface.
 */
export function instanceOfProfileTypeKeywords(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileTypeKeywordsFromJSON(json: any): ProfileTypeKeywords {
    return ProfileTypeKeywordsFromJSONTyped(json, false);
}

export function ProfileTypeKeywordsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileTypeKeywords {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'keyword': !exists(json, 'keyword') ? undefined : ((json['keyword'] as Array<any>).map(KeywordTypeFromJSON)),
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function ProfileTypeKeywordsToJSON(value?: ProfileTypeKeywords | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'keyword': value.keyword === undefined ? undefined : ((value.keyword as Array<any>).map(KeywordTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

