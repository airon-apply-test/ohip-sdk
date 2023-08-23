/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ConversionCodeMappingType } from './ConversionCodeMappingType';
import {
    ConversionCodeMappingTypeFromJSON,
    ConversionCodeMappingTypeFromJSONTyped,
    ConversionCodeMappingTypeToJSON,
} from './ConversionCodeMappingType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { IntegrationSystemType } from './IntegrationSystemType';
import {
    IntegrationSystemTypeFromJSON,
    IntegrationSystemTypeFromJSONTyped,
    IntegrationSystemTypeToJSON,
} from './IntegrationSystemType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PutConversionCodeMappingsRequest
 */
export interface PutConversionCodeMappingsRequest {
    /**
     * 
     * @type {IntegrationSystemType}
     * @memberof PutConversionCodeMappingsRequest
     */
    integrationSystem?: IntegrationSystemType;
    /**
     * List of Conversion Code Mappings.
     * @type {Array<ConversionCodeMappingType>}
     * @memberof PutConversionCodeMappingsRequest
     */
    conversionCodeMappings?: Array<ConversionCodeMappingType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PutConversionCodeMappingsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutConversionCodeMappingsRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PutConversionCodeMappingsRequest interface.
 */
export function instanceOfPutConversionCodeMappingsRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PutConversionCodeMappingsRequestFromJSON(json: any): PutConversionCodeMappingsRequest {
    return PutConversionCodeMappingsRequestFromJSONTyped(json, false);
}

export function PutConversionCodeMappingsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutConversionCodeMappingsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'integrationSystem': !exists(json, 'integrationSystem') ? undefined : IntegrationSystemTypeFromJSON(json['integrationSystem']),
        'conversionCodeMappings': !exists(json, 'conversionCodeMappings') ? undefined : ((json['conversionCodeMappings'] as Array<any>).map(ConversionCodeMappingTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PutConversionCodeMappingsRequestToJSON(value?: PutConversionCodeMappingsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'integrationSystem': IntegrationSystemTypeToJSON(value.integrationSystem),
        'conversionCodeMappings': value.conversionCodeMappings === undefined ? undefined : ((value.conversionCodeMappings as Array<any>).map(ConversionCodeMappingTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

