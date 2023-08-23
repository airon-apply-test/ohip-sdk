"use strict";
/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Cashiering API
 * APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.CompAccountingJournalToJSON = exports.CompAccountingJournalFromJSONTyped = exports.CompAccountingJournalFromJSON = exports.instanceOfCompAccountingJournal = void 0;
const runtime_1 = require("../runtime");
const FinancialPostingsType_1 = require("./FinancialPostingsType");
const InstanceLink_1 = require("./InstanceLink");
const WarningType_1 = require("./WarningType");
/**
 * Check if a given object implements the CompAccountingJournal interface.
 */
function instanceOfCompAccountingJournal(value) {
    let isInstance = true;
    return isInstance;
}
exports.instanceOfCompAccountingJournal = instanceOfCompAccountingJournal;
function CompAccountingJournalFromJSON(json) {
    return CompAccountingJournalFromJSONTyped(json, false);
}
exports.CompAccountingJournalFromJSON = CompAccountingJournalFromJSON;
function CompAccountingJournalFromJSONTyped(json, ignoreDiscriminator) {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        'journalPostings': !(0, runtime_1.exists)(json, 'journalPostings') ? undefined : (0, FinancialPostingsType_1.FinancialPostingsTypeFromJSON)(json['journalPostings']),
        'totalPages': !(0, runtime_1.exists)(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !(0, runtime_1.exists)(json, 'offset') ? undefined : json['offset'],
        'limit': !(0, runtime_1.exists)(json, 'limit') ? undefined : json['limit'],
        'hasMore': !(0, runtime_1.exists)(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !(0, runtime_1.exists)(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !(0, runtime_1.exists)(json, 'count') ? undefined : json['count'],
        'links': !(0, runtime_1.exists)(json, 'links') ? undefined : (json['links'].map(InstanceLink_1.InstanceLinkFromJSON)),
        'warnings': !(0, runtime_1.exists)(json, 'warnings') ? undefined : (json['warnings'].map(WarningType_1.WarningTypeFromJSON)),
    };
}
exports.CompAccountingJournalFromJSONTyped = CompAccountingJournalFromJSONTyped;
function CompAccountingJournalToJSON(value) {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        'journalPostings': (0, FinancialPostingsType_1.FinancialPostingsTypeToJSON)(value.journalPostings),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
        'links': value.links === undefined ? undefined : (value.links.map(InstanceLink_1.InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : (value.warnings.map(WarningType_1.WarningTypeToJSON)),
    };
}
exports.CompAccountingJournalToJSON = CompAccountingJournalToJSON;
