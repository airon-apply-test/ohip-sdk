/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Content Service
 * Opera Content Service offers capability to manage large content such as images and files.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { EmailRecipientType } from './EmailRecipientType';
import {
    EmailRecipientTypeFromJSON,
    EmailRecipientTypeFromJSONTyped,
    EmailRecipientTypeToJSON,
} from './EmailRecipientType';
import type { FolioReportCriteriaType } from './FolioReportCriteriaType';
import {
    FolioReportCriteriaTypeFromJSON,
    FolioReportCriteriaTypeFromJSONTyped,
    FolioReportCriteriaTypeToJSON,
} from './FolioReportCriteriaType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request to generate and email a folio report.
 * @export
 * @interface FolioReportToEmail
 */
export interface FolioReportToEmail {
    /**
     * 
     * @type {FolioReportCriteriaType}
     * @memberof FolioReportToEmail
     */
    folioInfo?: FolioReportCriteriaType;
    /**
     * 
     * @type {Array<EmailRecipientType>}
     * @memberof FolioReportToEmail
     */
    emailRecipients?: Array<EmailRecipientType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof FolioReportToEmail
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof FolioReportToEmail
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the FolioReportToEmail interface.
 */
export function instanceOfFolioReportToEmail(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FolioReportToEmailFromJSON(json: any): FolioReportToEmail {
    return FolioReportToEmailFromJSONTyped(json, false);
}

export function FolioReportToEmailFromJSONTyped(json: any, ignoreDiscriminator: boolean): FolioReportToEmail {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'folioInfo': !exists(json, 'folioInfo') ? undefined : FolioReportCriteriaTypeFromJSON(json['folioInfo']),
        'emailRecipients': !exists(json, 'emailRecipients') ? undefined : ((json['emailRecipients'] as Array<any>).map(EmailRecipientTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function FolioReportToEmailToJSON(value?: FolioReportToEmail | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'folioInfo': FolioReportCriteriaTypeToJSON(value.folioInfo),
        'emailRecipients': value.emailRecipients === undefined ? undefined : ((value.emailRecipients as Array<any>).map(EmailRecipientTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

